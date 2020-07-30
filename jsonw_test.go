package jsonw_test

import (
	"os"

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
