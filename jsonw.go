// Package jsonw offers a "quality-of-life" writer for dynamic JSON.
package jsonw

import (
	"bufio"
	"encoding/json"
	"io"
	"strconv"
)

// Writer writes JSON values to an output stream.
type Writer struct {
	depth int   // nesting level
	comma bool  // comma before value/name?
	state state // next expected token
	err   error // first error

	w *bufio.Writer
}

type state byte

const (
	item  state = iota // a top-level value or array item
	name               // a name in name/value pair
	value              // a value in name/value pair
)

// New returns a new JSON Writer that writes to w.
func New(w io.Writer) *Writer {
	return &Writer{w: bufio.NewWriter(w)}
}

// Object writes an object (a set of name/value pairs) to the stream.
// Writes within f will be nested in the object.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (j *Writer) Object(f func()) error {
	j.startValue()
	j.w.WriteByte('{')
	s := j.state
	j.state = name
	j.comma = false
	j.depth++
	f()
	j.depth--
	j.state = s
	j.comma = true
	j.w.WriteByte('}')
	j.endValue()
	return j.err
}

// Array writes an array (an ordered collection of values) to the stream.
// Writes within f will be nested in the array.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (j *Writer) Array(f func()) error {
	j.startValue()
	j.w.WriteByte('[')
	s := j.state
	j.state = item
	j.comma = false
	j.depth++
	f()
	j.depth--
	j.state = s
	j.comma = true
	j.w.WriteByte(']')
	j.endValue()
	return j.err
}

// Value writes a value to the stream.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (j *Writer) Value(v interface{}) error {
	j.startValue()
	buf, err := json.Marshal(v)
	if j.err == nil {
		j.err = err
	}
	j.w.Write(buf)
	j.endValue()
	return j.err
}

// Int writes an int value to the stream.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (j *Writer) Int(i int) error {
	j.startValue()
	j.w.WriteString(strconv.Itoa(i))
	j.endValue()
	return j.err
}

// Name writes a name to the stream.
// Returns this Writer, so you can fluently add the value.
// Panics if a value is expected.
func (j *Writer) Name(n string) *Writer {
	if j.state != name {
		panic("expected a value, not a name")
	}
	if j.comma {
		j.w.WriteByte(',')
	} else {
		j.comma = true
	}
	j.writeString(n)
	j.w.WriteByte(':')
	j.state = value
	return j
}

func (j *Writer) startValue() {
	if j.state == name {
		panic("expected a name, not a value")
	}
	if j.state == value {
		return
	}
	if j.comma {
		j.w.WriteByte(',')
	} else {
		j.comma = true
	}
}

func (j *Writer) endValue() {
	if j.state == name {
		panic("expected a name, not a value")
	}
	if j.state == value {
		j.state = name
		return
	}
	if j.depth == 0 {
		j.comma = false
		j.w.WriteByte('\n')
		err := j.w.Flush()
		if j.err == nil {
			j.err = err
		}
	}
}

func (j *Writer) writeString(s string) {
	for i := 0; i < len(s); i++ {
		if c := s[i]; false ||
			c < ' ' || c > '~' || // not printable ASCII
			c == '"' || c == '\\' || // need escape (JSON)
			c == '<' || c == '>' || c == '&' { // need escape (HTML/XML)

			// slow path
			buf, err := json.Marshal(s)
			if err != nil {
				panic(err)
			}
			j.w.Write(buf)
			return
		}
	}

	// fast path
	j.w.WriteByte('"')
	j.w.WriteString(s)
	j.w.WriteByte('"')
}
