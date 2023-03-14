package day23

//*割绳子：长度为n绳子，割成m端求最大成绩
func cuttingRope(n int) int {
	dp2 := make([]int, n+1) //*第二层动态规划数组，仅考虑第一段绳子j
	//*分成i个位置割开，记录分成两段后的最大值
	for i := 2; i <= n; i++ {
		//*遍历所有可能的第一段绳子长度j,记录最大值
		for j := 1; j < i; j++ {
			//*计算乘积j*max(i-j, dp[i-j])
			prod := j * max(i-j, dp2[i-j]) //*计算i-j
			dp2[i] = max(dp2[i], prod)     //*更新dp2[i]
		}
	}

	//?第一层动态规划数组，表示长度i的绳子割成若干段后乘积的最大值
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		//*遍历所有可能的第一段长度j
		for j := 1; j < i; j++ {
			prod := j * dp[i-j]
			dp[i] = max3(dp[i], prod, dp2[i])
		}
	}
	return dp[n]
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func max3(a, b, c int) int {
	tmp := 0
	if a > b {
		tmp = a
	} else {
		tmp = b
	}

	if tmp > c {
		return tmp
	}
	return c
}
