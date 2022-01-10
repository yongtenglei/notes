# 跳跃游戏 II [jumpGame II](https://leetcode-cn.com/problems/jump-game-ii/)

当位置在最大位置时再次起跳, 不遍历最后一个元素, 需要在倒数第二个元素就确定是否可以到达最后一个元素, 避免在最后一个元素进行[不必要的起跳].

```go

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	steps := 0
	end := 0
	max := 0

	for i := 0; i < len(nums)-1; i++ {
		nextMax := i + nums[i]

		max = IntMax(max, nextMax)

		if i == end {
			end = max
			steps++
		}

	}
	return steps
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}

	return b

}

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2, 3, 0, 1, 4}
	nums := []int{1, 2, 3}
	fmt.Println(jump(nums))
}
```
