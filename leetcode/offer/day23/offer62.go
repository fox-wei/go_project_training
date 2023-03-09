package day23

//*圆圈中剩下的数字：约瑟夫环问题
func lastRemaining(n int, m int) int {
	//约瑟夫环问题：f(n, m) = [f(n-1, m)+m] % n
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}
