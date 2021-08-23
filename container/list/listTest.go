package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	myList.PushFront(1)
	myList.PushFront(2)
	myList.PushFront(3)
	for myList.Len() > 0 {
		fmt.Println(myList.Remove(myList.Front()))
	}

	myList.PushFront(1)
	myList.PushFront(2)
	myList.PushFront(3)

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
