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
	flip bool
}

func parseArgs() (args Args) {
	flag.Usage = func() {
		log.Printf("usage: %s [options] <layout.toml> <footprints.toml>...", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.StringVar(&args.outputFilename, "output", "", "filename to write output SVG to")
	flag.BoolVar(&args.individualWires, "individual-wires", false, "if given, produce a separate image for each wire (rather than showing all at once); the argument to -output is considered to be the directory in which to place these images")
	flag.BoolVar(&args.flip, "flip", false, "if given, flip the image left-to-right, so that it shows the underside of the board")
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
	render(w, l, l.Wires, fps, args.flip)
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

	sort.Sort(byLength(l.Wires))

	for i := range l.Wires {
		filename := filepath.Join(args.outputFilename, fmt.Sprintf("wire%04d.svg", i))
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		w := bufio.NewWriter(f)
		render(w, l, l.Wires[i:i+1], fps, args.flip)
		w.Flush()
		f.Close()
	}
}

// Sort wires by increasing length
type byLength []*layout.Wire

func (wires byLength) Len() int {
	return len(wires)
}

func (wires byLength) Less(i, j int) bool {
	return wires[i].Length < wires[j].Length
}

func (wires byLength) Swap(i, j int) {
	wires[i], wires[j] = wires[j], wires[i]
}
