package main

import (
	"container/ring"
	"fmt"
)

func main() {
	ra := ring.New(3)

	fmt.Println("ra len=", ra.Len())

	for i := 0; i < ra.Len(); i++ {
		ra.Value = i
		ra = ra.Next()
	}

	ra.Do(f)
	fmt.Println()
}
func f(ab interface{}) {
	fmt.Print(ab, ",")
}
