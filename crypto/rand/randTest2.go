package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {

	// launches 2 generatores and the fanIn collector function
	c := fanIn2(genrt(), genrt())
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func fanIn2(a <-chan int, b <-chan int) <-chan string {
	c := make(chan string)
	// launch collector from a to channel
	go func() {
		var count int
		for {
			count += <-a
			c <- fmt.Sprintf("Tally of A is: %d", count)
		}
	}()
	// launch collector from b to channel
	go func() {
		var count int
		for {
			count += <-b
			c <- fmt.Sprintf("Tally of B is: %d", count)
		}
	}()

	return c
}

func genrt() <-chan int {
	c := make(chan int)
	// launch generator of Dice rolls
	go func() {
		for i := 0; ; i++ {
			dice, err := rand.Int(rand.Reader, big.NewInt(300))
			if err != nil {
				fmt.Println(err)
			}
			c <- int(dice.Int64()) + 1
			time.Sleep(time.Duration(1 * time.Millisecond))
		}
	}()
	return c
}
