package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func main(){
	c := make(chan os.Signal, 0)
	signal.Notify(c)
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}

func f1(){
	fmt.Println(os.Hostname())
	fmt.Println(strings.Join(os.Environ(),","))
	fmt.Println(os.Getenv("TMP"))
	fmt.Println(os.Getenv("aa"))
	//defer fmt.Println("aaaaaa")
	//os.Exit(0)
	fmt.Println(os.Getpid())
	//os.Exit(11111)
	fmt.Println(os.Stat("D:/goutil"))
	fmt.Println(os.Getwd())
	//os.Mkdir("D:/goutil/goutil/os/111",666)
	//os.MkdirAll("D:/goutil/goutil/os/222/3333",666)
	os.Remove("D:/goutil/goutil/os/111")
	fmt.Println(os.TempDir())
	fmt.Println(os.Args)
}
