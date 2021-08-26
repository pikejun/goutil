package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	t := hex.EncodeToString([]byte{100, 200, 255})
	fmt.Println(t)
	v, _ := hex.DecodeString("64c8ff")
	fmt.Println(v)
}
