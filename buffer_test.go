package jsonw_test

import (
	"fmt"
	"testing"

	"github.com/ncruces/go-jsonw"
)

func ExampleBuffer() {
	var jb jsonw.Buffer
	jb.Object(func() {
		jb.Name("ID").Int(1)
		jb.Name("Name").Value("Reds")
		jb.Name("Colors").Values("Crimson", "Red", "Ruby", "Maroon")
	})
	fmt.Print(jb.String())
	// Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

func TestBuffer_String(t *testing.T) {
	var jb jsonw.Buffer
	if got := jb.String(); got != "" {
		t.Errorf("got: %q", got)
	}

	jb.Int(1)
	if got := jb.String(); got != "1" {
		t.Errorf("got: %q", got)
	}

	jb.Int(1)
	jb.Value(true)
	if got := jb.String(); got != "true" {
		t.Errorf("got: %q", got)
	}
}
