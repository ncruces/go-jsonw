// Package jsonw offers a "quality-of-life" writer for dynamic JSON.
package jsonw

import (
	"encoding/json"
	"io"
	"strconv"
)

type jsonw struct {
	depth int   // nesting level
	comma bool  // comma before value/name?
	state state // next expected token
	err   error // first error
}

type writer interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

type flusher interface {
	Flush() error
}

type state byte

const (
	item  state = iota // a top-level value or array item
	name               // a name in name/value pair
	value              // a value in name/value pair
)

func (j *jsonw) object(w writer, f func()) error {
	j.startValue(w)
	w.WriteByte('{')
	s := j.state
	j.state = name
	j.comma = false
	j.depth++
	f()
	j.depth--
	j.state = s
	j.comma = true
	w.WriteByte('}')
	j.endValue(w)
	return j.err
}

func (j *jsonw) array(w writer, f func()) error {
	j.startValue(w)
	w.WriteByte('[')
	s := j.state
	j.state = item
	j.comma = false
	j.depth++
	f()
	j.depth--
	j.state = s
	j.comma = true
	w.WriteByte(']')
	j.endValue(w)
	return j.err
}

func (j *jsonw) value(w writer, v interface{}) error {
	j.startValue(w)
	buf, err := json.Marshal(v)
	if j.err == nil {
		j.err = err
	}
	w.Write(buf)
	j.endValue(w)
	return j.err
}

func (j *jsonw) int(w writer, i int) error {
	j.startValue(w)
	w.WriteString(strconv.Itoa(i))
	j.endValue(w)
	return j.err
}

func (j *jsonw) name(w writer, n string) {
	if j.state != name {
		panic("expected a value, not a name")
	}
	if j.comma {
		w.WriteByte(',')
	} else {
		j.comma = true
	}
	j.writeString(w, n)
	w.WriteByte(':')
	j.state = value
}

func (j *jsonw) startValue(w writer) {
	if j.state == name {
		panic("expected a name, not a value")
	}
	if j.state == value {
		return
	}
	if j.comma {
		w.WriteByte(',')
	} else {
		j.comma = true
	}
}

func (j *jsonw) endValue(w writer) {
	if j.state == name {
		panic("expected a name, not a value")
	}
	if j.state == value {
		j.state = name
		return
	}
	if j.depth == 0 {
		j.comma = false
		w.WriteByte('\n')
		if f, ok := w.(flusher); ok {
			err := f.Flush()
			if j.err == nil {
				j.err = err
			}
		}
	}
}

func (j *jsonw) writeString(w writer, s string) {
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
			w.Write(buf)
			return
		}
	}

	// fast path
	w.WriteByte('"')
	w.WriteString(s)
	w.WriteByte('"')
}
