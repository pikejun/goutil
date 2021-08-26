package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	h := sha512.New()
	io.WriteString(h, "12312312312312")
	fmt.Printf("%x\n", h.Sum(nil))
	h1 := sha256.New()
	io.WriteString(h1, "12312312312312")
	fmt.Printf("%x\n", h1.Sum(nil))
	h2 := sha1.New()
	io.WriteString(h2, "12312312312312")
	fmt.Printf("%x\n", h2.Sum(nil))
}
