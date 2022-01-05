package main

// canJump 贪心算法
func canJump(nums []int) bool {
	len := len(nums)
	if len == 1 {
		return true
	}

	max := 0

	for i := 0; i <= max; i++ {
		nextMax := nums[i] + i

		max = intMax(max, nextMax)

		if max >= len-1 {
			return true
		}
	}

	return false
}

func intMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// canJump2 动态规划
func canJump2(nums []int) bool {
	// bp[i] 代表第i个位置是否可以到达
	bp := make([]bool, len(nums), len(nums))

	bp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if bp[j] && j+nums[j] >= i {
				bp[i] = true
			}
		}
	}

	return bp[len(nums)-1]
}

func main() {

	arr := []int{2, 3, 1, 1, 4}
	print(canJump2(arr))
}
