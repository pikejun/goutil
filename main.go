package main

import (
	"fmt"
	"regexp"
)

func main() {
	//b,_:=regexp.Match("H.* ", []byte("Hello World!"))
	reg, err := regexp.Compile(`o\s+W.+`)
	fmt.Printf("%q,%v\n", reg.FindString("HelloWorld!"), err)
// true
}