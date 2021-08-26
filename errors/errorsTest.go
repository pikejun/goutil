package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("1111")
	fmt.Println(err1.Error())
}
