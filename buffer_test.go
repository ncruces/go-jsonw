package jsonw_test

import (
	"fmt"

	"github.com/ncruces/go-jsonw"
)

func ExampleBuffer() {
	var jb jsonw.Buffer
	jb.Object(func() {
		jb.Name("ID").Int(1)
		jb.Name("Name").Value("Reds")
		jb.Name("Colors").Value([]string{"Crimson", "Red", "Ruby", "Maroon"})
	})
	fmt.Print(jb.String())
	// Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}
