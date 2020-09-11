package jsonw_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ncruces/go-jsonw"
)

func Example() {
	jw := jsonw.New(os.Stdout)
	jw.Object(func() {
		jw.Name("ID").Value(1)
		jw.Name("Name").Value("Reds")
		jw.Name("Colors").Value([]string{"Crimson", "Red", "Ruby", "Maroon"})
	})
	// Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

func Benchmark(b *testing.B) {
	jw := jsonw.New(ioutil.Discard)
	for n := 0; n < b.N; n++ {
		jw.Object(func() {
			jw.Name("ID").Int(1)
			jw.Name("Name").Value("Reds")
			jw.Name("Colors").Value([]string{"Crimson", "Red", "Ruby", "Maroon"})
		})
	}
}
