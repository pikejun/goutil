package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {

	message := "12345678"
	i := len([]byte(message)) * 2
	buff := make([]byte, i)
	n := ascii85.Encode(buff, []byte(message))
	fmt.Println(string(buff[:n]))
	j, _, _ := ascii85.Decode(buff, buff, true)
	b := buff[:j]
	fmt.Println(string(b))
}
