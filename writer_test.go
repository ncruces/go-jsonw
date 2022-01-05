package jsonw_test

import (
	"io"
	"os"
	"testing"

	"github.com/ncruces/go-jsonw"
)

func ExampleWriter() {
	jw := jsonw.New(os.Stdout)
	jw.Object(func() {
		jw.Name("ID").Int(1)
		jw.Name("Name").Value("Reds")
		jw.Name("Colors").Value([]string{"Crimson", "Red", "Ruby", "Maroon"})
	})
	// Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

func BenchmarkWriter(b *testing.B) {
	jw := jsonw.New(io.Discard)
	for n := 0; n < b.N; n++ {
		jw.Object(func() {
			jw.Name("ID").Int(1)
			jw.Name("Name").Value("Reds")
			jw.Name("Colors").Value([]string{"Crimson", "Red", "Ruby", "Maroon"})
		})
	}
}
