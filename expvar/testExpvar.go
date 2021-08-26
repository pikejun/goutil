package main

import (
	. "expvar"
	"fmt"
)

func main() {
	var a String
	var b Int
	var c Map
	var d Func
	c.Set("abc", &a)
	a.Set("999")
	b.Set(1111)
	b.Add(1)
	d = t
	fmt.Println(a.Value())
	fmt.Println(b.Value())
	fmt.Println(c.Get("abc"))
	fmt.Println(d.String())
}
func t() interface{} {
	var a2 a
	a2.Age = 10
	a2.Name = "1111"
	return a2
}

type a struct {
	Name string
	Age  int
}

func (a1 *a) A() String {

	var t String
	t.Set("aaaa")
	return t
}

func (a1 *a) B() String {

	var t String
	t.Set("bbbb")
	return t
}
