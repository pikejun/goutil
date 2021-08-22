package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D:/400/1_maizhi2.txt"
	b, _ := os.Open(fileName)
	a := bufio.NewReader(b)
	defer b.Close()
	for {
		s, error := a.ReadString(byte('\n'))
		fmt.Println(s)
		if error != nil {
			break
		}
	}

	filename2 := "D:/400/1111.txt"
	os.Create(filename2)
	f, _ := os.OpenFile(filename2, os.O_WRONLY, 0)
	fw := bufio.NewWriter(f)
	defer f.Close()
	b.Seek(0, 0)
	for {
		s, error := a.ReadString(byte('\n'))
		fmt.Println(s)
		fw.Write([]byte("aaaaaaaaaaaaa" + s))
		if error != nil {
			break
		}
	}
}
