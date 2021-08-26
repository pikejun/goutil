package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	h := sha1.New()
	io.WriteString(h, "His money is twice tainted:")
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("%x", h.Sum(nil))
}
