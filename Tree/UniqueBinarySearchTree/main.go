package main

func NumTreesV1(n int) int {
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

func NumTreesV2(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
