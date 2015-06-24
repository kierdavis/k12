package main

// Standard library imports
import (
	"bufio"
	"flag"
	"log"
	"os"
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
}

func parseArgs() (args Args) {
	flag.Usage = func() {
		log.Printf("usage: %s [options] <layout.toml> <footprints.toml>...", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.StringVar(&args.outputFilename, "output", "", "filename to write output SVG to")
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
	
	if args.outputFilename == "" {
		args.outputFilename = args.layoutFilename + ".svg"
	}
	
	f, err := os.Create(args.outputFilename)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()
	
	w := bufio.NewWriter(f)
	render(w, l, fps)
	w.Flush()
}
