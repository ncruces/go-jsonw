package jsonw

import (
	"bytes"
	"io"
)

// Buffer writes JSON values to a buffer.
// Buffer implements io.Reader, io.WriterTo, json.Marshaler and fmt.Stringer
// by reading back the written JSON.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	b bytes.Buffer
	jsonw
}

// Object writes an object (a set of name/value pairs) to the buffer.
// Writes within f will be nested in the object.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Object(f func()) error {
	return b.object(&b.b, f)
}

// Array writes an array (an ordered collection of values) to the buffer.
// Writes within f will be nested in the array.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Array(f func()) error {
	return b.array(&b.b, f)
}

// Value writes a value to the buffer.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Value(v interface{}) error {
	return b.value(&b.b, v)
}

// Int writes an int value to the buffer.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Int(i int) error {
	return b.int(&b.b, i)
}

// Name writes a name to the buffer.
// Returns this Buffer, so you can fluently add the value.
// Panics if a value is expected.
func (b *Buffer) Name(n string) *Buffer {
	b.name(&b.b, n)
	return b
}

// Read reads the contents of the buffer into p.
// Implements io.Reader.
// Panics if a value is being written.
func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.depth != 0 {
		panic("value is incomplete")
	}
	return b.b.Read(p)
}

// WriteTo writes the contents of the buffer into w.
// Implements io.WriterTo.
// Panics if a value is being written.
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	if b.depth != 0 {
		panic("value is incomplete")
	}
	return b.b.WriteTo(w)
}

// String returns the contents of the unread portion of the buffer as a string.
// Panics if a value is being written.
func (b *Buffer) String() string {
	if b.depth != 0 {
		panic("value is incomplete")
	}
	return b.b.String()
}

// MarshalJSON returns the contents of the unread portion of the buffer
// as the JSON encoding of b.
// If the buffer is empty, returns null.
// Panics if a value is being written.
func (b *Buffer) MarshalJSON() ([]byte, error) {
	if b.depth != 0 {
		panic("value is incomplete")
	}
	if b.b.Len() == 0 {
		return []byte("null"), nil
	}
	return b.b.Bytes(), nil
}

// Reset resets the buffer to be empty.
func (b *Buffer) Reset() {
	b.jsonw = jsonw{}
	b.b.Reset()
}
