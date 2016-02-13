package util

// Standard library imports
import (
	"encoding/hex"
	"fmt"
	"image/color"
)

type Color struct {
	color.RGBA
}

func (c Color) MarshalText() (b []byte, err error) {
	b = make([]byte, 8)
	hex.Encode(b, []byte{c.RGBA.R, c.RGBA.G, c.RGBA.B, c.RGBA.A})
	return b, nil
}

func (c *Color) UnmarshalText(b []byte) (err error) {
	s := string(b)
	if s == "none" {
		c.RGBA.R = 0x00
		c.RGBA.G = 0x00
		c.RGBA.B = 0x00
		c.RGBA.A = 0x00
		return nil
	}
	var buf [4]byte
	switch len(b) {
	case 6:
		_, err = hex.Decode(buf[:3], b)
		buf[3] = 0xFF
	case 8:
		_, err = hex.Decode(buf[:], b)
	default:
		return fmt.Errorf("Color.UnmarshalText: expected a hexadecimal string in the form RRGGBB or RRGGBBAA, or \"none\" (not %q)", s)
	}
	if err != nil {
		return fmt.Errorf("Color.UnmarshalText: expected a hexadecimal string in the form RRGGBB or RRGGBBAA, or \"none\" (not %q)", s)
	}
	c.RGBA.R = buf[0]
	c.RGBA.G = buf[1]
	c.RGBA.B = buf[2]
	c.RGBA.A = buf[3]
	return nil
}

type Graphic struct {
	Path        string  `toml:"path"`
	Stroke      Color   `toml:"stroke"`
	StrokeWidth float64 `toml:"stroke_width"`
	Fill        Color   `toml:"fill"`
}
