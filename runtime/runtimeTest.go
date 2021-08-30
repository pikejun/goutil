package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go say("11111111")
	say("2222222")
}
func say(a string){
	for {
		runtime.Gosched()
		fmt.Println(a)
	}
}