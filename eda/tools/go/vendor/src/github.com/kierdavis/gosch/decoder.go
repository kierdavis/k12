package gosch

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

type Decoder struct {
	r      *bufio.Reader
	lineno int
}

func NewDecoder(r io.Reader) (d *Decoder) {
	return &Decoder{
		r:      bufio.NewReader(r),
		lineno: 0,
	}
}

func (d *Decoder) ReadLine() (line string, err error) {
	line, err = d.r.ReadString('\n')
	if err != nil {
		return "", err
	}
	d.lineno++
	return strings.TrimRight(line, "\r\n"), nil
}

func (d *Decoder) Decode() (f *File, err error) {
	f = new(File)
	err = d.decodeVersionRecord(f)
	if err != nil {
		return nil, err
	}

	f.Objects, err = d.decodeObjects()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (d *Decoder) decodeVersionRecord(f *File) (err error) {
	line, err := d.ReadLine()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("unexpected EOF when reading version record")
		}
		return err
	}

	parts := strings.SplitN(line, " ", 2)
	recordType := parts[0]
	recordBody := ""
	if len(parts) > 1 {
		recordBody = parts[1]
	}

	if recordType != "v" {
		return fmt.Errorf("at line %d: expected version ('v') record, not '%s'", d.lineno, recordType)
	}

	_, err = fmt.Sscanf(recordBody, "%s %d", &f.GedaVersion, &f.FormatVersion)
	if err != nil {
		return fmt.Errorf("at line %d: invalid record syntax: %s", d.lineno, err)
	}

	return nil
}

func (d *Decoder) decodeObjects() (objects []Object, err error) {
	var object Object

	for {
		line, err := d.ReadLine()
		if err != nil {
			if err == io.EOF {
				return objects, nil
			}
			return nil, err
		}

		parts := strings.SplitN(line, " ", 2)
		recordType := parts[0]
		recordBody := ""
		if len(parts) > 1 {
			recordBody = parts[1]
		}

		switch recordType {
		case "L":
			object = new(Line)
		case "G":
			object = new(Picture)
		case "B":
			object = new(Box)
		case "V":
			object = new(Circle)
		case "A":
			object = new(Arc)
		case "T":
			object = new(Text)
		case "N":
			object = new(Net)
		case "U":
			object = new(Bus)
		case "P":
			object = new(Pin)
		case "C":
			object = new(Component)
		case "H":
			object = new(Path)

		case "{":
			if object == nil {
				return nil, fmt.Errorf("at line %d: attribute block with no parent object", d.lineno)
			}

			// apply attributes to previous object
			err = d.decodeAttributes(object)
			if err != nil {
				return nil, err
			}

			continue

		default:
			return nil, fmt.Errorf("at line %d: invalid record type '%s'", d.lineno, recordType)
		}

		err = object.Decode(d, recordBody)
		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}

	return objects, nil
}

func (d *Decoder) decodeAttributes(object Object) (err error) {
	for {
		line, err := d.ReadLine()
		if err != nil {
			if err == io.EOF {
				return fmt.Errorf("at line %d: unterminated attribute block", d.lineno)
			}
			return err
		}

		parts := strings.SplitN(line, " ", 2)
		recordType := parts[0]
		recordBody := ""
		if len(parts) > 1 {
			recordBody = parts[1]
		}

		var text Text

		switch recordType {
		case "T": // pass
		case "}":
			return nil // end
		default:
			return fmt.Errorf("at line %d: invalid record type '%s' in attribute block", d.lineno, recordType)
		}

		err = text.Decode(d, recordBody)
		if err != nil {
			return err
		}

		parts = strings.SplitN(text.Text, "=", 2)
		attr := Attribute{
			Name:  parts[0],
			Value: parts[1],
			Text:  text,
		}

		object.AddAttr(attr)
	}
}

func (x *Line) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d %d %d %d %d", &x.Pos1.X,
		&x.Pos1.Y, &x.Pos2.X, &x.Pos2.Y, &x.Color, &x.LineParams.Width,
		&x.LineParams.CapStyle, &x.LineParams.DashStyle, &x.LineParams.DashLength,
		&x.LineParams.DashSpace)

	if err != nil {
		return fmt.Errorf("at line %d: invalid line ('L') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Picture) Decode(d *Decoder, recordBody string) (err error) {
	var mirrored, embedded rune

	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %c %c", &x.Pos.X, &x.Pos.Y,
		&x.Width, &x.Height, &x.Angle, &mirrored, &embedded)

	if err != nil {
		return fmt.Errorf("at line %d: invalid picture ('G') record syntax: %s", d.lineno, err)
	}

	x.Mirrored = mirrored == '1'
	x.Embedded = embedded == '1'

	x.Filename, err = d.ReadLine()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("at line %d: incomplete picture ('G') record", d.lineno)
		}
		return err
	}

	if x.Embedded {
		var buf []byte

		for {
			line, err := d.ReadLine()
			if err != nil {
				if err == io.EOF {
					return fmt.Errorf("at line %d: incomplete picture ('G') record", d.lineno)
				}
				return err
			}

			if line == "." {
				break
			}

			chunk, err := base64.StdEncoding.DecodeString(line)
			if err != nil {
				return fmt.Errorf("at line %d: invalid base64 syntax: %s", d.lineno, err)
			}

			buf = append(buf, chunk...)
		}

		x.EmbeddedData = buf
	}

	return nil
}

func (x *Box) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
		&x.Pos.X, &x.Pos.Y, &x.Width, &x.Height, &x.Color, &x.LineParams.Width,
		&x.LineParams.CapStyle, &x.LineParams.DashStyle, &x.LineParams.DashLength,
		&x.LineParams.DashSpace, &x.FillParams.Type, &x.FillParams.LineWidth,
		&x.FillParams.Angle1, &x.FillParams.Pitch1, &x.FillParams.Angle2,
		&x.FillParams.Pitch2)

	if err != nil {
		return fmt.Errorf("at line %d: invalid box ('B') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Circle) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
		&x.Pos.X, &x.Pos.Y, &x.Radius, &x.Color, &x.LineParams.Width,
		&x.LineParams.CapStyle, &x.LineParams.DashStyle, &x.LineParams.DashLength,
		&x.LineParams.DashSpace, &x.FillParams.Type, &x.FillParams.LineWidth,
		&x.FillParams.Angle1, &x.FillParams.Pitch1, &x.FillParams.Angle2,
		&x.FillParams.Pitch2)

	if err != nil {
		return fmt.Errorf("at line %d: invalid circle ('V') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Arc) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d %d %d %d %d %d", &x.Pos.X,
		&x.Pos.Y, &x.Radius, &x.StartAngle, &x.SweepAngle, &x.Color,
		&x.LineParams.Width, &x.LineParams.CapStyle, &x.LineParams.DashStyle,
		&x.LineParams.DashLength, &x.LineParams.DashSpace)

	if err != nil {
		return fmt.Errorf("at line %d: invalid arc ('A') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Text) Decode(d *Decoder, recordBody string) (err error) {
	var visible rune
	var numLines int

	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %c %d %d %d %d", &x.Pos.X,
		&x.Pos.Y, &x.Color, &x.Size, &visible, &x.AttrVisibility, &x.Angle,
		&x.Alignment, &numLines)

	if err != nil {
		return fmt.Errorf("at line %d: invalid text ('T') record syntax: %s", d.lineno, err)
	}

	x.Visible = visible == '1'

	text := ""
	for i := 0; i < numLines; i++ {
		line, err := d.ReadLine()
		if err != nil {
			if err == io.EOF {
				return fmt.Errorf("at line %d: incomplete text ('T') record", d.lineno)
			}
			return err
		}
		text += line + "\n"
	}

	// Remove trailing newline
	x.Text = text[:len(text)-1]

	return nil
}

func (x *Net) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d", &x.Pos1.X, &x.Pos1.Y,
		&x.Pos2.X, &x.Pos2.Y, &x.Color)

	if err != nil {
		return fmt.Errorf("at line %d: invalid net ('N') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Bus) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d", &x.Pos1.X, &x.Pos1.Y,
		&x.Pos2.X, &x.Pos2.Y, &x.Color, &x.RipperDir)

	if err != nil {
		return fmt.Errorf("at line %d: invalid bus ('U') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Pin) Decode(d *Decoder, recordBody string) (err error) {
	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d %d", &x.Pos1.X, &x.Pos1.Y,
		&x.Pos2.X, &x.Pos2.Y, &x.Color, &x.Type, &x.ActivePoint)

	if err != nil {
		return fmt.Errorf("at line %d: invalid pin ('P') record syntax: %s", d.lineno, err)
	}

	return nil
}

func (x *Component) Decode(d *Decoder, recordBody string) (err error) {
	var selectable, mirrored rune

	_, err = fmt.Sscanf(recordBody, "%d %d %c %d %c %s", &x.Pos.X, &x.Pos.Y,
		&selectable, &x.Angle, &mirrored, &x.SymName)

	if err != nil {
		return fmt.Errorf("at line %d: invalid component ('C') record syntax: %s", d.lineno, err)
	}

	x.Selectable = selectable == '1'
	x.Mirrored = mirrored == '1'

	return nil
}

func (x *Path) Decode(d *Decoder, recordBody string) (err error) {
	var numLines int

	_, err = fmt.Sscanf(recordBody, "%d %d %d %d %d %d %d %d %d %d %d %d %d",
		&x.Color, &x.LineParams.Width, &x.LineParams.CapStyle, &x.LineParams.DashStyle,
		&x.LineParams.DashLength, &x.LineParams.DashSpace, &x.FillParams.Type,
		&x.FillParams.LineWidth, &x.FillParams.Angle1, &x.FillParams.Pitch1,
		&x.FillParams.Angle2, &x.FillParams.Pitch2, &numLines)

	if err != nil {
		return fmt.Errorf("at line %d: invalid path ('H') record syntax: %s", d.lineno, err)
	}

	data := ""
	for i := 0; i < numLines; i++ {
		line, err := d.ReadLine()
		if err != nil {
			if err == io.EOF {
				return fmt.Errorf("at line %d: incomplete path ('H') record", d.lineno)
			}
			return err
		}
		data += line + " "
	}

	// Remove trailing space
	x.Data = PathData(data[:len(data)-1])

	return nil
}
