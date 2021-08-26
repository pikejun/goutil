package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed data
var f embed.FS

func main() {
	fmt.Println(s)
	b, _ := f.ReadFile("data/g1/aaa.yml")
	fmt.Println(string(b))
}
