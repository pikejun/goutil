package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

// 这是一个基础的使用用例，创建一个编码器，对数据进行编码，然后使用解码器接收数据
func main() {
	// 初始化编码器，创建一个decoder实例
	var network bytes.Buffer        // 标准输入
	enc := gob.NewEncoder(&network) // 编码
	dec := gob.NewDecoder(&network) // 解码
	// 编码器发送数据
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// 解码器接收数据
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}
