package offer24

//*栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	//?使用辅助栈，栈顶元素不能与出站序列的第一元素，进行入栈；相等则弹出
	if len(popped) == 0 || len(pushed) == 0 || len(popped) != len(pushed) {
		return false
	}

	stack := []int{} //*辅助栈
	i := 0           //*弹出序列指针
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && stack[0] == popped[i] {
			stack = stack[1:]
			i++
		}
	}
	return len(stack) == 0
}
