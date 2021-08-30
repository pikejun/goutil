package main

import "fmt"

func  main()  {

	fmt.Printf("%s %[1]T %3d %10x,%3.3f %v %T\n","aaaa",10123131,11,1234.12,31,23123)
	fmt.Printf("%q %x %#x %v","abc","abc","abc",map[float64]int{1: 1})
}
