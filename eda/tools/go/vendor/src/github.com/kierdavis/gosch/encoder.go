package gosch

import (
	"encoding/base64"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) Encoder {
	return Encoder{w}
}

func (e Encoder) Encode(f *File) (err error) {
	_, err = fmt.Fprintf(e.w, "v %s %d\n", f.GedaVersion, f.FormatVersion)
	if err != nil {
		return err
	}

	for _, object := range f.Objects {
		err = e.EncodeObjectWithAttrs(object)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e Encoder) EncodeObjectWithAttrs(o Object) (err error) {
	err = o.Encode(e)
	if err != nil {
		return err
	}

	attrs := o.Attrs()
	if len(attrs) > 0 {
		err = e.EncodeAttrs(attrs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e Encoder) EncodeAttrs(attrs []Attribute) (err error) {
	_, err = fmt.Fprint(e.w, "{\n")
	if err != nil {
		return err
	}

	for _, attr := range attrs {
		attr.Text.Text = fmt.Sprintf("%s=%s", attr.Name, attr.Value)

		// Can attributes have their own attributes??
		// If so, replace with call to EncodeObjectWithAttrs
		err = attr.Text.Encode(e)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(e.w, "}\n")
	return err
}

func (x *Line) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "L %d %d %d %d %d %s\n", x.Pos1.X, x.Pos1.Y, x.Pos2.X, x.Pos2.Y, x.Color, x.LineParams)
	return err
}

func (x *Picture) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "G %d %d %d %d %d %c %c\n%s\n", x.Pos.X, x.Pos.Y, x.Width, x.Height, x.Angle, bool2char(x.Mirrored), bool2char(x.Embedded), x.Filename)
	if err != nil {
		return err
	}

	if x.Embedded {
		encodedData := base64.StdEncoding.EncodeToString(x.EmbeddedData)

		lineLength := 64
		pos := 0
		for pos < len(encodedData) {
			end := pos + lineLength
			if end > len(encodedData) {
				end = len(encodedData)
			}
			_, err = fmt.Fprintf(e.w, "%s\n", encodedData[pos:end])
			if err != nil {
				return err
			}
			pos = end
		}

		fmt.Fprint(e.w, ".\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *Box) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "B %d %d %d %d %d %s %s\n", x.Pos.X, x.Pos.Y, x.Width, x.Height, x.Color, x.LineParams, x.FillParams)
	return err
}

func (x *Circle) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "V %d %d %d %d %s %s\n", x.Pos.X, x.Pos.Y, x.Radius, x.Color, x.LineParams, x.FillParams)
	return err
}

func (x *Arc) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "A %d %d %d %d %d %d %s\n", x.Pos.X, x.Pos.Y, x.Radius, x.StartAngle, x.SweepAngle, x.Color, x.LineParams)
	return err
}

// A backslash (escaped both for Go and for the regexp) followed by a character
// that's not an underscore or backslash.
var invalidBackslash = regexp.MustCompile("\\\\[^_\\\\]")

func replInvalidBackslash(s string) string {
	// s is two chars, the backslash and the following character
	// Prepend a backslash so that it contains two backslashes in a row, escaping
	// the backslash and leaving the following character intact.
	return "\\" + s
}

func (x *Text) Encode(e Encoder) (err error) {
	text := invalidBackslash.ReplaceAllStringFunc(x.Text, replInvalidBackslash)
	numLines := 1 + strings.Count(text, "\n")

	_, err = fmt.Fprintf(e.w, "T %d %d %d %d %c %d %d %d %d\n%s\n", x.Pos.X, x.Pos.Y, x.Color, x.Size, bool2char(x.Visible), x.AttrVisibility, x.Angle, x.Alignment, numLines, text)
	return err
}

func (x *Net) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "N %d %d %d %d %d\n", x.Pos1.X, x.Pos1.Y, x.Pos2.X, x.Pos2.Y, x.Color)
	return err
}

func (x *Bus) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "U %d %d %d %d %d %d\n", x.Pos1.X, x.Pos1.Y, x.Pos2.X, x.Pos2.Y, x.Color, x.RipperDir)
	return err
}

func (x *Pin) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "P %d %d %d %d %d %d %d\n", x.Pos1.X, x.Pos1.Y, x.Pos2.X, x.Pos2.Y, x.Color, x.Type, x.ActivePoint)
	return err
}

func (x *Component) Encode(e Encoder) (err error) {
	_, err = fmt.Fprintf(e.w, "C %d %d %c %d %c %s\n", x.Pos.X, x.Pos.Y, bool2char(x.Selectable), x.Angle, bool2char(x.Mirrored), x.SymName)
	if err != nil {
		return err
	}

	if len(x.EmbeddedObjects) > 0 {
		_, err = fmt.Fprint(e.w, "[\n")
		if err != nil {
			return err
		}

		for _, o := range x.EmbeddedObjects {
			err = e.EncodeObjectWithAttrs(o)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Fprint(e.w, "]\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *Path) Encode(e Encoder) (err error) {
	data := string(x.Data)

	// break data into lines
	var lines []string
	lineLength := 80
	for len(data) > lineLength {
		pos := lineLength
		for pos >= 0 && data[pos] != ' ' {
			pos--
		}
		lines = append(lines, data[:pos])
		data = data[pos+1:]
	}
	if len(data) > 0 {
		lines = append(lines, data)
	}

	_, err = fmt.Fprintf(e.w, "H %d %s %s %d\n", x.Color, x.LineParams, x.FillParams, len(lines))
	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err = fmt.Fprintf(e.w, "%s\n", line)
		if err != nil {
			return err
		}
	}

	return nil
}
