package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:age`
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	p1 := &Person{}

	//{"name":"aa","address":"chengdu","Age":19}
	json.Unmarshal([]byte(`{"name":"aa","address":"chengdu","Age":19}`), p1)

	fmt.Println(p1)
}
