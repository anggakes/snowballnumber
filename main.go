package main

import (
	"fmt"
)

func main() {
	// you can write to stdout for debugging purposes, e.g.
	//fmt.Println("This is a debug message")
	//fmt.Println("numcpu", runtime.NumCPU())
	test := 9
	//k := SnowballNumberSimple(test)
	//fmt.Println(len(k))
	//fmt.Println(k)
	l := SnowballNumberReduceIteration(test)
	fmt.Println(len(l))
	fmt.Println(l)
	//m := SnowballNumberArithmetic(test)
	//fmt.Println(len(m))
	//fmt.Println(m)
	//
	//n := SnowballNumberWithAritmatikChannel(test)
	//fmt.Println(len(n))
	//fmt.Println(n)
}
