package main

// Standard library imports
import (
	"fmt"
	"image/color"
	"io"
)

// Local imports
import (
	"format/footprint"
	"format/layout"
	"format/util"
)

func render(w io.Writer, l *layout.Layout, fps map[string]*footprint.Footprint) {
	minX := l.MinX - 1
	minY := l.MinY - 1
	maxX := l.MaxX + 1
	maxY := l.MaxY + 1
	
	innerWidth := maxX - minX
	innerHeight := maxY - minY
	
	outerHeight := 600
	outerWidth := (outerHeight * innerWidth) / innerHeight
	
	fmt.Fprint(w, "<?xml version='1.0' standalone='no'?>\n")
	fmt.Fprint(w, "<!DOCTYPE svg PUBLIC \"-//W3C//DTD SVG 1.1//EN\" \"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd\">\n")
	fmt.Fprintf(w, "<svg width='%d' height='%d' viewBox='%d %d %d %d' version='1.1' xmlns='http://www.w3.org/2000/svg'>\n", outerWidth, outerHeight, minX, minY, innerWidth, innerHeight)
	//fmt.Fprintf(w, "  <rect x='%d' y='%d' width='%d' height='%d' stroke='#999900' stroke-width='0.1' fill='none'/>\n", l.MinX, l.MinY, innerWidth-2, innerHeight-2)
	
	renderGraphics(w, l)
	renderComponents(w, l, fps)
	renderWires(w, l, fps)
	
	fmt.Fprint(w, "</svg>\n")
}

func renderGraphics(w io.Writer, l *layout.Layout) {
	fmt.Fprintf(w, "  <g id='graphics'>\n")
	
	for _, graphic := range l.Graphics {
		renderGraphic(w, graphic)
	}
	
	fmt.Fprintf(w, "  </g>\n")
}

func renderComponents(w io.Writer, l *layout.Layout, fps map[string]*footprint.Footprint) {
	for refdes, comp := range l.Components {
		fp := fps[comp.Footprint]
		if fp == nil {
			panic("invalid footprint")
		}
		
		fmt.Fprintf(w, "  <g id='comp-%s' transform='translate(%d, %d) rotate(%d)'>\n", refdes, comp.X, comp.Y, comp.Rotate*90)
		
		for _, coords := range fp.PinCoords {
			fmt.Fprintf(w, "    <circle cx='%d' cy='%d' r='0.4' stroke='none' fill='green'/>\n", coords[0], coords[1])
			fmt.Fprintf(w, "    <circle cx='%d' cy='%d' r='0.2' stroke='none' fill='white'/>\n", coords[0], coords[1])
		}
		
		for _, graphic := range fp.Graphics {
			renderGraphic(w, graphic)
		}
		
		fmt.Fprintf(w, "    <text x='%f' y='%f' text-anchor='middle' dominant-baseline='central' font-size='1px'>%s</text>\n", comp.LabelX, comp.LabelY, refdes)
		fmt.Fprint(w, "  </g>\n")
	}
}

func renderWires(w io.Writer, l *layout.Layout, fps map[string]*footprint.Footprint) {
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
		
		c := util.Color{color.RGBA{0x00, 0x00, 0xFF, 0xFF}}
		n := l.Nets[wire.Net]
		if n != nil {
			c = n.Color
		}
		
		fmt.Fprintf(w, "  <line x1='%d' y1='%d' x2='%d' y2='%d' stroke='%s' stroke-width='0.2'/>\n", x1, y1, x2, y2, colorStr(c))
	}
}

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

func renderGraphic(w io.Writer, graphic *util.Graphic) {
	fmt.Fprintf(w, "    <path d='%s' stroke='%s' stroke-width='%f' fill='%s'/>\n", graphic.Path, colorStr(graphic.Stroke), graphic.StrokeWidth, colorStr(graphic.Fill))
}

func colorStr(c util.Color) (s string) {
	s = "none"
	if c.RGBA.A == 0xFF {
		s = fmt.Sprintf("#%02X%02X%02X", c.RGBA.R, c.RGBA.G, c.RGBA.B)
	}
	return s
}
