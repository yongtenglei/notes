package main

import "fmt"

func generate(numRows int) [][]int {
	nums := make([][]int, numRows, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		nums[i] = row
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = nums[i-1][j] + nums[i-1][j-1]
		}
	}

	return nums

}

func getRow(rowIndex int) []int {
	rowIndex = rowIndex + 1
	nums := make([][]int, rowIndex, rowIndex)

	for i := 0; i < rowIndex; i++ {
		row := make([]int, i+1)
		nums[i] = row
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = nums[i-1][j] + nums[i-1][j-1]
		}
	}

	return nums[rowIndex-1]

}

func main() {
	//nums := make([][]int, 2, 2)

	//a := make([]int, 2)

	//nums[0] = a

	//fmt.Println(nums)

	//a[0] = 1

	//fmt.Println(nums)

	//a[1] = 2

	//fmt.Println(nums)

	fmt.Println(generate(5))
	fmt.Println(getRow(3))
}
