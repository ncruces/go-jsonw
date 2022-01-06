package jsonw

import (
	"bytes"
)

// Buffer writes JSON values to a buffer.
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
	if b.depth == 0 {
		b.Reset()
	}
	return b.object(&b.b, f)
}

// Array writes an array (an ordered collection of values) to the buffer.
// Writes within f will be nested in the array.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Array(f func()) error {
	if b.depth == 0 {
		b.Reset()
	}
	return b.array(&b.b, f)
}

// Value writes a value to the buffer.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Value(v interface{}) error {
	if b.depth == 0 {
		b.Reset()
	}
	return b.value(&b.b, v)
}

// Values writes an array (an ordered collection) of values to the buffer.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Values(v ...interface{}) error {
	if b.depth == 0 {
		b.Reset()
	}
	return b.value(&b.b, v)
}

// Int writes an int value to the buffer.
// Returns the first serialization error found.
// Panics if a name is expected.
func (b *Buffer) Int(i int) error {
	if b.depth == 0 {
		b.Reset()
	}
	return b.int(&b.b, i)
}

// Name writes a name to the buffer.
// Returns this Buffer, so you can fluently add the value.
// Panics if a value is expected.
func (b *Buffer) Name(n string) *Buffer {
	b.name(&b.b, n)
	return b
}

// String returns the last written JSON value as a string.
// Panics if the value has not been fully written.
func (b *Buffer) String() string {
	if b.depth != 0 {
		panic("value not fully written")
	}
	return b.b.String()
}

// String returns the last written JSON value as a byte slice.
// The returned slice is valid only until the next buffer modification.
// Panics if the value has not been fully written.
func (b *Buffer) Bytes() []byte {
	if b.depth != 0 {
		panic("value not fully written")
	}
	return b.b.Bytes()
}

// MarshalJSON returns the last written JSON value as the JSON encoding of b.
// The returned slice is valid only until the next buffer modification.
// If the buffer is empty, returns null.
// Panics if the value has not been fully written.
func (b *Buffer) MarshalJSON() ([]byte, error) {
	if b.depth != 0 {
		panic("value not fully written")
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
