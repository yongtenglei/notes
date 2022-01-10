# 生成括号 [generateParenthesis](https://leetcode-cn.com/problems/generate-parentheses/)

## DFS

要点:

1. 长度到达 2\*N 时, 返回结果.

2. 左括号数量小于 N 时, 可以添加左括号.

3. 右括号数量小于左括号数量时, 添加右括号.

```go
func generateParenthesis(n int) []string {
	res := new([]string)

	dfs(0, 0, n, "", res)

	return *res

}

func dfs(lnum, rnum int, n int, s string, res *[]string) *[]string {
	if len(s) == 2*n {
		*res = append(*res, s)
		return res
	}

	if lnum < n {
		dfs(lnum+1, rnum, n, s+"(", res)
	}

	if rnum < lnum {
		dfs(lnum, rnum+1, n, s+")", res)
	}
	return nil
}

```
