# 跳跃游戏 [jump Game](https://leetcode-cn.com/problems/jump-game/)

## 动态规划

dp[i] 代表 i 处是否可以到达, 最后是否到达即判断 dp 是否为 true.

```go

func canJump(nums []int) bool {
	bp := make([]bool, len(nums), len(nums))

  // 第一步已经到达
	bp[0] = true

  // 对于每一个位置i, 确定是否可以到达:
  // 如果前面的某个位置j可以到达, 并且此位置+可以跳跃的步数>i, 则i可以到达.
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if bp[j] && j+nums[j] >= i {
				bp[i] = true
			}
		}
	}

	return bp[len(nums)-1]
}

```

## 贪心算法

不关心每一个位置是否可以走到, 只关心可以到达的最后位置.

```go
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	// 可以走到的最大位置
	max := 0

	for i := 0; i <= max; i++ {
		nextPositon := i + nums[i]

		// 更新最大位置
		max = intMax(nextPositon, max)

		// 可以到达的最大位置大与最后的位置, 则退出.
		// 否则继续尝试max之前的位置, 更新最大位置.
		if max >= len(nums)-1 {
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

```
