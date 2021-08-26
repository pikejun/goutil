package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, x := range []int64{-1, 1, 2, 4, 127, 128, 255, 256, 512} {
		n := binary.PutVarint(buf, x)
		fmt.Print(x, "输出的可变长度为：", n, "，十六进制为：")
		fmt.Printf("%x\n", buf[:n])
	}
}

func main2() {
	buf := new(bytes.Buffer)
	var pi int32 = 676 //2*16*16+10*16+4
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes()) // 18 2d 44 54 fb 21 09 40
}
