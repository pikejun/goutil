package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	message = "12345678"
)

func Md5(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return sha
}

func main() {

	fmt.Println(Md5(message))
}
