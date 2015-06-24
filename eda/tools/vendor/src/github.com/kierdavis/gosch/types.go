package gosch

import (
	"fmt"
)

//go:generate stringer -output=types_string.go -type=Color,CapStyle,DashStyle,FillType,AttrVisibility,TextAlignment,PinType,PinPoint

// A number of mils (1000ths of an inch), used to specify lengths and coordinates.
type Mils int

// An angle in degrees.
type Degrees int

// A coordinate within the schematic/symbol.
// Positive x axis extends to the right, negative to the left.
// Positive y axis extends upwards, negative downwards.
type Coord struct {
	X Mils
	Y Mils
}

// An index into the color palette.
type Color uint

// Standard color palette.
// See: http://wiki.geda-project.org/geda:file_format_spec, section "Colors".
const (
	BackgroundColor        Color = 0
	PinColor               Color = 1
	NetEndpointColor       Color = 2
	GraphicColor           Color = 3
	NetColor               Color = 4
	AttributeColor         Color = 5
	LogicBubbleColor       Color = 6
	DotsGridColor          Color = 7
	DetachedAttributeColor Color = 8
	TextColor              Color = 9
	BusColor               Color = 10
	SelectColor            Color = 11
	BoundingBoxColor       Color = 12
	ZoomBoxColor           Color = 13
	StrokeColor            Color = 14
	LockColor              Color = 15
	OutputBackgroundColor  Color = 16
	Freestyle1Color        Color = 17
	Freestyle2Color        Color = 18
	Freestyle3Color        Color = 19
	Freestyle4Color        Color = 20
	JunctionColor          Color = 21
	MeshGridMajorColor     Color = 22
	MeshGridMinorColor     Color = 23
)

// Line end (cap) style.
type CapStyle uint

const (
	NoCap     CapStyle = 0
	SquareCap CapStyle = 1
	RoundCap  CapStyle = 2
)

// Line dash style.
type DashStyle uint

const (
	Solid   DashStyle = 0
	Dotted  DashStyle = 1
	Dashed  DashStyle = 2
	Center  DashStyle = 3
	Phantom DashStyle = 4
)

// Box fill type.
type FillType uint

const (
	Hollow FillType = 0 // no fill
	Fill   FillType = 1 // solid color
	Mesh   FillType = 2
	Hatch  FillType = 3
	Void   FillType = 4
)

// Specifies which parts of an attribute text object are visible
type AttrVisibility uint

const (
	NameAndValue AttrVisibility = 0
	ValueOnly    AttrVisibility = 1
	NameOnly     AttrVisibility = 2
)

func (v AttrVisibility) NameVisible() bool {
	return v == NameAndValue || v == NameOnly
}

func (v AttrVisibility) ValueVisible() bool {
	return v == NameAndValue || v == ValueOnly
}

// Text alignment (anchor) point.
type TextAlignment uint

const (
	BottomLeft   TextAlignment = 0
	MiddleLeft   TextAlignment = 1
	TopLeft      TextAlignment = 2
	BottomCenter TextAlignment = 3
	MiddleCenter TextAlignment = 4
	TopCenter    TextAlignment = 5
	BottomRight  TextAlignment = 6
	MiddleRight  TextAlignment = 7
	TopRight     TextAlignment = 8
)

// Pin type.
type PinType uint

const (
	NormalPin PinType = 0
	BusPin    PinType = 1
)

// Pin end point number.
type PinPoint uint

const (
	Point1 PinPoint = 0
	Point2 PinPoint = 1
)

type LineParams struct {
	Width      Mils
	CapStyle   CapStyle
	DashStyle  DashStyle
	DashLength Mils // If DashStyle is Solid or Dotted, this field is ignored.
	DashSpace  Mils // If DashStyle is Solid, this field is ignored.
}

func (p LineParams) String() string {
	dashLength := p.DashLength
	dashSpace := p.DashSpace

	switch p.DashStyle {
	case Solid:
		dashLength = -1
		dashSpace = -1
	case Dotted:
		dashLength = -1
	}

	return fmt.Sprintf("%d %d %d %d %d", p.Width, p.CapStyle, p.DashStyle, dashLength, dashSpace)
}

type FillParams struct {
	Type      FillType
	LineWidth Mils
	Angle1    Degrees // Only used if FillType is Mesh or Hatch.
	Pitch1    Mils    // Only used if FillType is Mesh or Hatch.
	Angle2    Degrees // Only used if FillType is Mesh.
	Pitch2    Mils    // Only used if FillType is Mesh.
}

func (p FillParams) String() string {
	lineWidth := p.LineWidth
	angle1 := p.Angle1
	pitch1 := p.Pitch1
	angle2 := p.Angle2
	pitch2 := p.Pitch2

	switch p.Type {
	case Hollow, Fill:
		lineWidth = -1
		angle1 = -1
		pitch1 = -1
		angle2 = -1
		pitch2 = -1
	case Hatch:
		angle2 = -1
		pitch2 = -1
	}

	return fmt.Sprintf("%d %d %d %d %d %d", p.Type, lineWidth, angle1, pitch1, angle2, pitch2)
}

// Path data follows a subset of the SVG path data standard.
// See http://wiki.geda-project.org/geda:file_format_spec?&#path_data
// (or http://www.w3.org/TR/SVG/paths.html)
type PathData string

type Attribute struct {
	Name  string
	Value string
	Text
}

type File struct {
	GedaVersion   string
	FormatVersion int
	Objects       []Object
}

type Object interface {
	// Return a list of all attributes on the object.
	Attrs() []Attribute
	
	// Get the attribute with the given name.
	GetAttr(string) Attribute
	
	// Add an attribute with the given name and value, overwriting an existing
	// attribute with the same name if present.
	SetAttr(Attribute)
	
	// Add an attribute with the given name and value, creating a new attribute
	// even if one with the same name already exists.
	AddAttr(Attribute)

	Encode(Encoder) error
	Decode(*Decoder, string) error
}

type BaseObject struct {
	Attributes []Attribute
}

func (o *BaseObject) Attrs() []Attribute {
	return o.Attributes
}

func (o *BaseObject) GetAttr(name string) (attr Attribute) {
	for _, attr := range o.Attributes {
		if attr.Name == name {
			return attr
		}
	}
	return Attribute{}
}

func (o *BaseObject) SetAttr(attr Attribute) {
	for _, a := range o.Attributes {
		if a.Name == attr.Name {
			a.Value = attr.Value
			a.Text = attr.Text
			return
		}
	}

	o.Attributes = append(o.Attributes, attr)
}

func (o *BaseObject) AddAttr(attr Attribute) {
	o.Attributes = append(o.Attributes, attr)
}

type Line struct {
	BaseObject
	Pos1       Coord
	Pos2       Coord
	Color      Color
	LineParams LineParams
}

type Picture struct {
	BaseObject
	Pos          Coord // Lower left corner.
	Width        Mils
	Height       Mils
	Angle        Degrees // Only 0, 90, 180 and 270 allowed.
	Mirrored     bool
	Embedded     bool
	Filename     string
	EmbeddedData []byte
}

type Box struct {
	BaseObject
	Pos        Coord // Lower left corner.
	Width      Mils
	Height     Mils
	Color      Color
	LineParams LineParams
	FillParams FillParams
}

type Circle struct {
	BaseObject
	Pos        Coord // Center.
	Radius     Mils
	Color      Color
	LineParams LineParams
	FillParams FillParams
}

type Arc struct {
	BaseObject
	Pos        Coord // Center.
	Radius     Mils
	StartAngle Degrees
	SweepAngle Degrees
	Color      Color
	LineParams LineParams
}

type Text struct {
	BaseObject
	Pos            Coord // Position of anchor point, as determined by the Alignment field.
	Color          Color
	Size           int            // Font size in points (a point is 1/72 of an inch).
	Visible        bool           // Whether or not the text is visible.
	AttrVisibility AttrVisibility // Only used for attributes.
	Angle          int            // Only 0, 90, 180 and 270 allowed.
	Alignment      TextAlignment  // Anchor point.

	// Can contain newlines. Overbar text is denoted by surrounding it with "\_"
	// on each side (e.g. "\_CS\_", "R/\_W\_"). Backslashes can be escaped with
	// "\\" (although all backslashes not matching either of these two patterns
	// are automatically escaped).
	Text string
}

type Net struct {
	BaseObject
	Pos1  Coord
	Pos2  Coord
	Color Color
}

type Bus struct {
	BaseObject
	Pos1      Coord
	Pos2      Coord
	Color     Color
	RipperDir int // Direction of bus rippers (0, 1 or -1).
}

type Pin struct {
	BaseObject
	Pos1        Coord
	Pos2        Coord
	Color       Color
	Type        PinType
	ActivePoint PinPoint
}

type Component struct {
	BaseObject
	Pos             Coord
	Selectable      bool
	Angle           Degrees  // Only 0, 90, 180 and 270 allowed.
	Mirrored        bool     // True if mirrored around Y axis.
	SymName         string   // File name of symbol file (e.g. "7400-1.sym"), without the path.
	EmbeddedObjects []Object // If non-empty, the contents of this slice will be used as the embedded contents of this component.
}

type Path struct {
	BaseObject
	Color      Color
	LineParams LineParams
	FillParams FillParams
	Data       PathData
}
