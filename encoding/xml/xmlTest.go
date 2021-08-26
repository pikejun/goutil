package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name    string `xml:"name"`
	Address string `xml:"address"`
	Age     int    `xml:"age"`
}

func main() {
	p := &Person{"gogo", "ddd", 19}
	b, _ := xml.Marshal(p)
	fmt.Println(string(b))
	//<Person><name>gogo</name><address>ddd</address><age>19</age></Person>
	xml.Unmarshal([]byte(`<Person><name>gogo2</name><address>ddd2</address><age>190</age></Person>`), p)
	fmt.Println(p)
}
