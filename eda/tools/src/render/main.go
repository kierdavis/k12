package main

// Standard library imports
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// Local imports
import (
	"format/footprint"
	"format/layout"
)

type Args struct {
	outputFilename string
	layoutFilename string
	footprintFilenames []string
	individualWires bool
}

func parseArgs() (args Args) {
	flag.Usage = func() {
		log.Printf("usage: %s [options] <layout.toml> <footprints.toml>...", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.StringVar(&args.outputFilename, "output", "", "filename to write output SVG to")
	flag.BoolVar(&args.individualWires, "individual-wires", false, "if given, produce a separate image for each wire (rather than showing all at once); the argument to -output is considered to be the directory in which to place these images")
	flag.Parse()
	
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}
	
	args.layoutFilename = flag.Arg(0)
	args.footprintFilenames = flag.Args()[1:]
	
	return args
}

func loadLayout(filename string) (l *layout.Layout) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()
	
	l, err = layout.Read(f)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	
	return l
}

func loadFootprints(fps map[string]*footprint.Footprint, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()
	
	fps1, err := footprint.Read(f)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	
	for name, fp := range fps1 {
		fps[name] = fp
	}
}

func main() {
	args := parseArgs()
	l := loadLayout(args.layoutFilename)
	
	fps := make(map[string]*footprint.Footprint)
	for _, filename := range args.footprintFilenames {
		loadFootprints(fps, filename)
	}
	
	if args.individualWires {
		doIndividualWiresRender(args, l, fps)
	} else {
		doNormalRender(args, l, fps)
	}
}

func doNormalRender(args Args, l *layout.Layout, fps map[string]*footprint.Footprint) {
	if args.outputFilename == "" {
		args.outputFilename = args.layoutFilename + ".svg"
	}
	
	f, err := os.Create(args.outputFilename)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()
	
	w := bufio.NewWriter(f)
	render(w, l, l.Wires, fps)
	w.Flush()
}

func doIndividualWiresRender(args Args, l *layout.Layout, fps map[string]*footprint.Footprint) {
	if args.outputFilename == "" {
		args.outputFilename = args.layoutFilename + "-wires"
	}

	err := os.MkdirAll(args.outputFilename, 0755)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	sort.Sort(wireSort{l, fps})

	for i := range l.Wires {
		filename := filepath.Join(args.outputFilename, fmt.Sprintf("wire%04d.svg", i))
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		w := bufio.NewWriter(f)
		render(w, l, l.Wires[i:i+1], fps)
		w.Flush()
		f.Close()
	}
}

type wireSort struct {
	l *layout.Layout
	fps map[string]*footprint.Footprint
}

func (w wireSort) Len() int {
	return len(w.l.Wires)
}

func (w wireSort) Less(i, j int) bool {
	return squaredWireLength(w.l.Wires[i], w.l, w.fps) < squaredWireLength(w.l.Wires[j], w.l, w.fps)
}

func (w wireSort) Swap(i, j int) {
	w.l.Wires[i], w.l.Wires[j] = w.l.Wires[j], w.l.Wires[i]
}

func squaredWireLength(wire *layout.Wire, l *layout.Layout, fps map[string]*footprint.Footprint) int {
	c1 := l.Components[wire.Component1]
	if c1 == nil {
		log.Fatalf("undefined component '%s'", wire.Component1)
	}
	x1, y1 := pinPos(c1, fps, wire.Pin1)

	c2 := l.Components[wire.Component2]
	if c2 == nil {
		log.Fatalf("undefined component '%s'", wire.Component2)
	}
	x2, y2 := pinPos(c2, fps, wire.Pin2)

	dx := x2 - x1
	dy := y2 - y1
	return dx*dx + dy*dy
}


