package main

// Standard library imports
import (
	"flag"
	"log"
	"io"
	"os"
)

// Local imports
import (
	"format/footprint"
	"format/layout"
	"format/netlist"
)

type Args struct {
	outputFilename string
	layoutFilename string
	netlistFilename string
	footprintFilenames []string
}

func parseArgs() (args Args) {
	flag.Usage = func() {
		log.Printf("usage: %s [options] <layout.toml> <netlist.net> <footprints.toml>...", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.StringVar(&args.outputFilename, "output", "", "file to write generated layout to")
	flag.Parse()
	
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(2)
	}
	
	args.layoutFilename = flag.Arg(0)
	args.netlistFilename = flag.Arg(1)
	args.footprintFilenames = flag.Args()[2:]
	
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

func loadNetlist(filename string) (nets netlist.Netlist) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()
	
	nets, err = netlist.Read(f)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	
	return nets
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
	nets := loadNetlist(args.netlistFilename)
	
	fps := make(map[string]*footprint.Footprint)
	for _, filename := range args.footprintFilenames {
		loadFootprints(fps, filename)
	}
	
	Map(l, fps, nets)
	
	var w io.Writer
	
	if args.outputFilename == "" {
		w = os.Stdout
	} else {
		f, err := os.Create(args.outputFilename)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		defer f.Close()
		w = f
	}
	
	layout.Write(w, l)
	
	totalInches := measure(l, fps) / 10.0
	totalMetres := totalInches * 0.0254
	log.Printf("routing complete using %.1f inches (%.2f m) of wire", totalInches, totalMetres)
}
