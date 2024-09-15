package main

func NumTrees(n int) int {
	dp := make([]*int, n+1)
	return solve(n, dp)
}

func solve(n int, dp []*int) int {
	if n <= 1 {
		return 1
	}

	if dp[n] != nil {
		return *dp[n]
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += solve(i-1, dp) * solve(n-i, dp)
	}
	return ans
}
