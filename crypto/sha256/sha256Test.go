package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	var a string = "12312312"
	h := sha256.New()
	io.WriteString(h, a)
	fmt.Printf("%x\n", h.Sum(nil))

	h2 := sha1.New()
	io.WriteString(h2, a)
	fmt.Printf("%x\n", h2.Sum(nil))
}
