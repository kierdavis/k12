package footprint

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

type Footprint struct {
	PinCoords [][2]int        `toml:"pin_coords"`
	Graphics  []*util.Graphic `toml:"graphics"`
}

func Read(r io.Reader) (footprints map[string]*Footprint, err error) {
	_, err = toml.DecodeReader(r, &footprints)
	if err != nil {
		return nil, err
	}
	return footprints, nil
}

func Write(w io.Writer, footprints map[string]*Footprint) (err error) {
	return toml.NewEncoder(w).Encode(footprints)
}
