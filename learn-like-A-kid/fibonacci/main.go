package main

import (
	"fmt"
	"time"
)

// Fibonacci implemented with definition
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciWithNB implemented with definition as well as having a map to store values which are calculated before
func FibonacciWithNB(n int, nb map[int]int) int {
	if v, ok := nb[n]; ok {
		return v
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	v := FibonacciWithNB(n-1, nb) + FibonacciWithNB(n-2, nb)

	nb[n] = v

	return v
}

// FibonacciIter implemented in iteration way
func FibonacciIter(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

// FibonacciIter implemented in iteration way and return Fibonacci Array
func FibonacciIterArray(n int) []int {
	array := []int{0}
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
		array = append(array, a)
	}

	return array
}

func main() {
	//fmt.Println("======Fibonacci=======")
	//s := time.Now()
	//a := Fibonacci(50)
	//e := time.Now().Sub(s)
	//fmt.Println(a, e)

	//fmt.Println("======FibonacciWithNB=======")
	//s := time.Now()
	//b := FibonacciWithNB(5, map[int]int{})
	//e := time.Now().Sub(s)
	//fmt.Println(b, e)

	//fmt.Println("======FibonacciIter=======")
	//s = time.Now()
	//c := FibonacciIter(1000)
	//e = time.Now().Sub(s)
	//fmt.Println(c, e)

	fmt.Println("======FibonacciIterArray=======")
	s := time.Now()
	c := FibonacciIterArray(20)
	e := time.Now().Sub(s)
	fmt.Println(c, e)
}
