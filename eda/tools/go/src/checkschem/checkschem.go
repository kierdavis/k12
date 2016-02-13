package main

// Standard library imports
import (
	"flag"
	"log"
	"os"
	"strings"
)

// Vendored imports
import (
	"github.com/kierdavis/gosch"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Printf("usage: %s [options] schematic.sch...", os.Args[0])
		os.Exit(2)
	}

	schematics := loadSchematics(flag.Args()...)

	check(schematics)
}

func loadSchematics(filenames ...string) (schematics map[string]*gosch.File) {
	schematics = make(map[string]*gosch.File)

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

		schematics[filename] = schematic
	}

	return schematics
}

func check(schematics map[string]*gosch.File) {
	inputs := make(map[string]bool)
	outputs := make(map[string]bool)

	for filename, schematic := range schematics {
		for _, objectInterface := range schematic.Objects {
			switch object := objectInterface.(type) {
			case *gosch.Component:
				if object.SymName == "input-2.sym" || object.SymName == "output-2.sym" || object.SymName == "io-1.sym" {
					netStr := object.GetAttr("net").Value
					valueStr := object.GetAttr("value").Value

					parts := strings.Split(netStr, ":")
					if len(parts) != 2 || parts[1] != "1" {
						log.Printf("%s: '%s' symbol with value '%s': invalid syntax for 'net' attribute", filename, object.SymName, valueStr)
						continue
					}

					net := parts[0]
					expectedValue := net
					if expectedValue[0] == '-' {
						expectedValue = "\\_" + expectedValue[1:] + "\\_"
					}
					if expectedValue != valueStr {
						log.Printf("%s: '%s' symbol with value '%s': 'net' and 'value' attributes do not correspond", filename, object.SymName, valueStr)
						continue
					}

					if object.SymName == "input-2.sym" {
						inputs[net] = true
					} else if object.SymName == "output-2.sym" {
						outputs[net] = true
					} else {
						inputs[net] = true
						outputs[net] = true
					}
				}
			}
		}
	}

	for net, _ := range inputs {
		if !outputs[net] {
			log.Printf("net '%s' occurs as an input but not as an output", net)
		}
	}

	for net, _ := range outputs {
		if !inputs[net] {
			log.Printf("net '%s' occurs as an output but not as an input", net)
		}
	}
}
