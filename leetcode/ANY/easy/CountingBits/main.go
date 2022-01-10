package main

import "fmt"

func countBits(n int) []int {
	res := make([]int, n+1, n+1)
	res[0] = 0

	for i := 1; i < n+1; i++ {
		counter := 0
		numInBinary := fmt.Sprintf("%b", i)
		for _, c := range numInBinary {
			if c == '1' {
				counter++
			}
		}

		res[i] = counter
	}

	return res
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
