package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

func main() {
	message := "12345678"
	dst := make([]byte, 100)

	base32.StdEncoding.Encode(dst, []byte(message))

	fmt.Println(string(dst))

	base64.StdEncoding.Encode(dst, []byte(message))

	fmt.Println(string(dst))
}
