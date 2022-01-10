package main

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
