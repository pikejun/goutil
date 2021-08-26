package main

import (
	"fmt"
	"math/big"
)
import . "crypto/elliptic"

func main() {
	p224 := P224()
	x, _ := new(big.Int).SetString("ae99feebb5d26945b54892092a8aee02912930fa41cd114e40447301", 16)
	y, _ := new(big.Int).SetString("482580a0ec5bc47e88bc8c378632cd196cb3fa058a7114eb03054c9", 16)
	if p224.IsOnCurve(x, y) {
		fmt.Println(x, y)
	}
	b := Marshal(p224, x, y)
	fmt.Println(b)
	x1, y1 := Unmarshal(p224, b)
	fmt.Println(x1, y1)
	if x1 != nil || y1 != nil {
		fmt.Println("FAIL: unmarshaling a point not on the curve succeeded")
	}
}
