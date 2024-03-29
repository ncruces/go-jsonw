package jsonw_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ncruces/go-jsonw"
)

func ExampleWriter() {
	jw := jsonw.New(os.Stdout)
	jw.Object(func() {
		jw.Name("ID").Int(1)
		jw.Name("Name").Value("Reds")
		jw.Name("Colors").Values("Crimson", "Red", "Ruby", "Maroon")
	})
	// Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

func TestWriter(t *testing.T) {
	var buf strings.Builder
	jw := jsonw.New(&buf)
	jw.Int(1)
	jw.Value(true)
	if got := buf.String(); got != "1\ntrue\n" {
		t.Errorf("got: %q", got)
	}
}

func BenchmarkWriter(b *testing.B) {
	jw := jsonw.New(io.Discard)
	for n := 0; n < b.N; n++ {
		jw.Object(func() {
			jw.Name("ID").Int(1)
			jw.Name("Name").Value("Reds")
			jw.Name("Colors").Values("Crimson", "Red", "Ruby", "Maroon")
		})
	}
}
