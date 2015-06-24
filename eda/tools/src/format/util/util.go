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
	if c.RGBA.A == 0x00 {
		return []byte("none"), nil
	}
	b = make([]byte, 6)
	hex.Encode(b, []byte{c.RGBA.R, c.RGBA.G, c.RGBA.B})
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
	if len(b) != 6 {
		return fmt.Errorf("Color.UnmarshalText: expected a 6-character hexadecimal string in the form RRGGBB, or \"none\" (not %q)", s)
	}
	var buf [3]byte
	_, err = hex.Decode(buf[:], b)
	if err != nil {
		return fmt.Errorf("Color.UnmarshalText: expected a 6-character hexadecimal string in the form RRGGBB, or \"none\" (not %q)", s)
	}
	c.RGBA.R = buf[0]
	c.RGBA.G = buf[1]
	c.RGBA.B = buf[2]
	c.RGBA.A = 0xFF
	return nil
}

type Graphic struct {
    Path string `toml:"path"`
    Stroke Color `toml:"stroke"`
    StrokeWidth float64 `toml:"stroke_width"`
    Fill Color `toml:"fill"`
}
