package day7

//*斐波那契数列实现
//?动态规划实现
func fib(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}

	return a
}

//?青蛙跳台阶问题：问题本质斐波那锲数列类似
//?n级跳台有f(n)种方法，最后一步的时候：跳1级或者跳2两级。
//?最后一步：1级:f(n-1);两级:f(n-2)=》f(n)=f(n-1)+f(n-2)；这是斐波那契数列数列的性质相同
func newWay(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1_000_000_007
	}
	return a
}
