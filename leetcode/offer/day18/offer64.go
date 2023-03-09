package day18

//*求1+2...+n
//*要求不能使用for，while,if,else,switch等关键字
//*思路：采用递归实现，但是要使用if关键字；这里可以用短路原则规避
func traver(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNums(n-1)
}
func sumNums(n int) int {
	//*短路原则
	_ = n > 0 && func() bool {
		n += sumNums(n - 1)
		return true
	}()
	return n
}

func sumNums1(n int) int {
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	return n
}
