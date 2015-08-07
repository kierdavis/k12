package layout

// Standard library imports
import (
	"io"
)

// Vendored imports
import (
	"github.com/BurntSushi/toml"
)

// Local imports
import (
	"format/util"
)

type Layout struct {
	Title string `toml:"title"`
	MinX int `toml:"min_x"`
	MinY int `toml:"min_y"`
	MaxX int `toml:"max_x"`
	MaxY int `toml:"max_y"`
	Components map[string]*Component `toml:"components"`
	Nets map[string]*Net `toml:"nets"`
	Wires []*Wire `toml:"wires"`
	Graphics []*util.Graphic `toml:"graphics"`
}

type Component struct {
	Footprint string `toml:"footprint"`
	X int `toml:"x"`
	Y int `toml:"y"`
	Rotate int `toml:"rotate"` // number of 90 degree clockwise rotations
	LabelX float64 `toml:"label_x"`
	LabelY float64 `toml:"label_y"`
}

type Net struct {
	Color util.Color `toml:"color"`
}

type Wire struct {
	Component1 string `toml:"component1"`
	Pin1 int `toml:"pin1"`
	Component2 string `toml:"component2"`
	Pin2 int `toml:"pin2"`
	Net string `toml:"net"`
}

func Read(r io.Reader) (layout *Layout, err error) {
	_, err = toml.DecodeReader(r, &layout)
	if err != nil {
		return nil, err
	}
	return layout, nil
}

func Write(w io.Writer, layout *Layout) (err error) {
	return toml.NewEncoder(w).Encode(layout)
}
