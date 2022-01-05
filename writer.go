package jsonw

import (
	"bufio"
	"io"
)

// Writer writes JSON values to an io.Writer.
type Writer struct {
	w *bufio.Writer
	jsonw
}

// New returns a new JSON Writer that writes to w.
func New(w io.Writer) *Writer {
	return &Writer{w: bufio.NewWriter(w)}
}

// Object writes an object (a set of name/value pairs) to the writer.
// Writes within f will be nested in the object.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (w *Writer) Object(f func()) error {
	return w.object(w.w, f)
}

// Array writes an array (an ordered collection of values) to the writer.
// Writes within f will be nested in the array.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (w *Writer) Array(f func()) error {
	return w.array(w.w, f)
}

// Value writes a value to the writer.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (w *Writer) Value(v interface{}) error {
	return w.value(w.w, v)
}

// Value writes a value to the writer.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (w *Writer) Values(v ...interface{}) error {
	return w.value(w.w, v)
}

// Int writes an int value to the writer.
// Returns the first serialization error, or (at top level) io error, found.
// Panics if a name is expected.
func (w *Writer) Int(i int) error {
	return w.int(w.w, i)
}

// Name writes a name to the writer.
// Returns this Writer, so you can fluently add the value.
// Panics if a value is expected.
func (w *Writer) Name(n string) *Writer {
	w.name(w.w, n)
	return w
}
