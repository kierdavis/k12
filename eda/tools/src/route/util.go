package main

// Standard library imports
import (
	"math"
)

// Local imports
import (
	"format/footprint"
	"format/layout"
)

func pinPos(c *layout.Component, fps map[string]*footprint.Footprint, pin int) (x, y int) {
	fp := fps[c.Footprint]
	if fp == nil {
		panic("invalid footprint")
	}
	pinCoords := fp.PinCoords[pin-1]
	x, y = pinCoords[0], pinCoords[1]
	for i := 0; i < c.Rotate; i++ {
		x, y = -y, x
	}
	x += c.X
	y += c.Y
	return x, y
}

// returns the total length of wire present in the layout, in units (tenths of an inch)
func measure(l *layout.Layout, fps map[string]*footprint.Footprint) (totalLength float64) {
	for _, wire := range l.Wires {
		c1 := l.Components[wire.Component1]
		if c1 == nil {
			panic("invalid component")
		}
		x1, y1 := pinPos(c1, fps, wire.Pin1)

		c2 := l.Components[wire.Component2]
		if c2 == nil {
			panic("invalid component")
		}
		x2, y2 := pinPos(c2, fps, wire.Pin2)

		dx, dy := x2-x1, y2-y1
		totalLength += math.Sqrt(float64(dx*dx + dy*dy))
	}

	return totalLength
}
