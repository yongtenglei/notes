# 杨辉三角 [PascalTriangle](https://leetcode-cn.com/problems/pascals-triangle/)

<div align=center><img src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif">
</div>

## 杨辉三角 I

```go
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

```

## 杨辉三角 [II](https://leetcode-cn.com/problems/pascals-triangle-ii/)

```go
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

```

## 另一种解法

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNly1gy34qhl6vdj30io0n041a.jpg">

</div>

## 收获

```go
func main() {
	nums := make([][]int, 2, 2)

	a := make([]int, 2)

	nums[0] = a

	fmt.Println(nums)

	a[0] = 1

	fmt.Println(nums)

	a[1] = 2

	fmt.Println(nums)

}

```

```go
// output
[[0 0] []]
[[1 0] []]
[[1 2] []]

```

切片的底层是数组
