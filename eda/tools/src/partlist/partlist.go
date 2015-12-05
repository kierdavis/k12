package main

// Standard library imports
import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

// Vendored imports
import (
	"github.com/kierdavis/gosch"
)

var symPathFlag = flag.String("sympath", "", "directories in which to search for symbols, in addition to /usr/share/gEDA/sym (colon-separated)")
var outputFileFlag = flag.String("output", "parts.txt", "file to write output to")

var symSearchPath = []string{"/usr/share/gEDA/sym"}
var symIndex = make(map[string]string)
var loadedSyms = make(map[string]map[string]string)

type Component struct {
	sym      string // name of symbol file
	attrs    map[string]string
	symAttrs map[string]string
}

func (c *Component) loadSymAttrs() {
	symAttrs, ok := loadedSyms[c.sym]
	if ok {
		//log.Printf("used cached copy of %s", c.sym)
		c.symAttrs = symAttrs
		return
	}

	symAttrs = make(map[string]string)
	c.symAttrs = symAttrs
	loadedSyms[c.sym] = symAttrs

	filename := symIndex[c.sym]
	if filename == "" {
		log.Printf("error: could not locate symbol '%s'", c.sym)
		return
	}

	//log.Printf("reading %s", filename)

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("error: failed to open %s: %s", filename, err)
		return
	}

	d := gosch.NewDecoder(f)
	symbol, err := d.Decode()
	f.Close()
	if err != nil {
		log.Printf("error: failed to parse %s: %s", filename, err)
		return
	}

	for _, object := range symbol.Objects {
		text, ok := object.(*gosch.Text)
		if ok {
			parts := strings.SplitN(text.Text, "=", 2)
			if len(parts) == 2 {
				name := parts[0]
				value := parts[1]
				symAttrs[name] = value
			}
		}
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Printf("usage: %s [options] schematic.sch...", os.Args[0])
		os.Exit(2)
	}

	symPath := *symPathFlag
	outputFile := *outputFileFlag

	symSearchPath = append(symSearchPath, strings.Split(symPath, ":")...)

	indexSyms()

	objects := loadSchematics(flag.Args()...)
	comps := process(objects)

	f, err := os.Create(outputFile)
	if err != nil {
		log.Printf("error: failed to open %s: %s", outputFile, err)
		os.Exit(1)
	}
	defer f.Close()

	w := tabwriter.NewWriter(f, 8, 8, 2, ' ', 0)

	printICs(w, comps)
	printConns(w, comps)
	printOther(w, comps)

	//log.Printf("wrote %s", outputFile)
}

func indexSyms() {
	for _, root := range symSearchPath {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Printf("error: failed to search %s: %s", path, err)
				return filepath.SkipDir
			}
			name := info.Name()
			if strings.HasSuffix(name, ".sym") {
				symIndex[name] = path
			}
			return nil
		})
		if err != nil {
			log.Printf("error: failed to search %s: %s", root, err)
		}
	}
}

func loadSchematics(filenames ...string) (objects []gosch.Object) {
	for _, filename := range filenames {
		//log.Printf("reading %s", filename)

		f, err := os.Open(filename)
		if err != nil {
			log.Printf("error: failed to open %s: %s", filename, err)
			continue
		}

		d := gosch.NewDecoder(f)
		schematic, err := d.Decode()
		f.Close()
		if err != nil {
			log.Printf("error: failed to parse %s: %s", filename, err)
			continue
		}

		objects = append(objects, schematic.Objects...)
	}

	return objects
}

func process(objects []gosch.Object) (comps map[string]*Component) {
	comps = make(map[string]*Component)

	for _, object := range objects {
		c, ok := object.(*gosch.Component)
		if ok {
			attrs := make(map[string]string)
			for _, attr := range object.Attrs() {
				attrs[attr.Name] = attr.Value
			}

			refdes := attrs["refdes"]
			if refdes != "" {
				if comps[refdes] != nil {
					for name, value := range comps[refdes].attrs {
						attrs[name] = value
					}
				}

				comps[refdes] = &Component{
					sym:   c.SymName,
					attrs: attrs,
				}
			}
		}
	}

	return comps
}

var icRefdesRegexp = regexp.MustCompile("U(\\d+)")

func printICs(w *tabwriter.Writer, comps map[string]*Component) {
	var keys sort.IntSlice

	for refdes := range comps {
		m := icRefdesRegexp.FindStringSubmatch(refdes)
		if m != nil {
			n, _ := strconv.ParseInt(m[1], 10, 0)
			keys = append(keys, int(n))
		}
	}

	if len(keys) > 0 {
		sort.Sort(keys)

		fmt.Fprintf(w, "##### ICs #####\n\n")
		w.Flush()

		fmt.Fprintf(w, "REFDES\tDEVICE\tFOOTPRINT\tAPPLICATION DESCRIPTION\tDEVICE DESCRIPTION\n")

		for _, n := range keys {
			refdes := fmt.Sprintf("U%d", n)
			comp := comps[refdes]
			comp.loadSymAttrs()

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", refdes, comp.attrs["device"], comp.attrs["footprint"], comp.attrs["x-appdesc"], comp.symAttrs["description"])

			delete(comps, refdes)
		}

		fmt.Fprintf(w, "\n\n")
		w.Flush()
	}
}

var connRefdesRegexp = regexp.MustCompile("CONN(\\d+)")

func printConns(w *tabwriter.Writer, comps map[string]*Component) {
	var keys sort.IntSlice

	for refdes := range comps {
		m := connRefdesRegexp.FindStringSubmatch(refdes)
		if m != nil {
			n, _ := strconv.ParseInt(m[1], 10, 0)
			keys = append(keys, int(n))
		}
	}

	if len(keys) > 0 {
		sort.Sort(keys)

		fmt.Fprintf(w, "##### Connectors #####\n\n")
		w.Flush()

		fmt.Fprintf(w, "REFDES\tDEVICE\tFOOTPRINT\tAPPLICATION DESCRIPTION\n")

		for _, n := range keys {
			refdes := fmt.Sprintf("CONN%d", n)
			comp := comps[refdes]
			//comp.loadSymAttrs()

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", refdes, comp.attrs["device"], comp.attrs["footprint"], comp.attrs["x-appdesc"])

			delete(comps, refdes)
		}

		fmt.Fprintf(w, "\n\n")
		w.Flush()
	}
}

func printOther(w *tabwriter.Writer, comps map[string]*Component) {
	var keys sort.StringSlice

	for refdes := range comps {
		keys = append(keys, refdes)
	}

	if len(keys) > 0 {
		sort.Sort(keys)

		fmt.Fprintf(w, "##### Other #####\n\n")
		w.Flush()

		fmt.Fprintf(w, "REFDES\tDEVICE\tFOOTPRINT\tAPPLICATION DESCRIPTION\n")

		for _, refdes := range keys {
			comp := comps[refdes]
			//comp.loadSymAttrs()

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", refdes, comp.attrs["device"], comp.attrs["footprint"], comp.attrs["x-appdesc"])

			//delete(comps, refdes)
		}

		fmt.Fprintf(w, "\n\n")
		w.Flush()
	}
}
