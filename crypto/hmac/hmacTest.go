package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	message = "hello world!"
	secret  = "0933e54e76b24731a2d84b6b463ec04c1"
)

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func main() {

	fmt.Println(ComputeHmacSha256(message, secret))
}
