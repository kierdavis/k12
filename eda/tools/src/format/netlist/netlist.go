package netlist

// Standard library imports
import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var whitespace = regexp.MustCompile("\\s+")

type Netlist map[string][]Node

type Node struct {
	Component string
	Pin       int
}

func Read(r io.Reader) (nets Netlist, err error) {
	nets = make(map[string][]Node)

	scanner := bufio.NewScanner(r)
	prefix := ""

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t\r\n")
		if len(line) == 0 {
			continue
		}

		if strings.HasSuffix(line, "\\") {
			prefix += line[:len(line)-1]
			continue
		}

		line = prefix + line
		prefix = ""

		parts := whitespace.Split(line, -1)
		net := parts[0]
		nodes := make([]Node, 0, len(parts)-1)

		for _, part := range parts[1:] {
			p := strings.IndexRune(part, '-')
			comp := part[:p]
			pin, err := strconv.ParseInt(part[p+1:], 10, 0)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, Node{comp, int(pin)})
		}

		nets[net] = nodes
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return nets, nil
}
