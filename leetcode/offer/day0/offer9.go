package main

//*offer.09用栈实现两个队列，队列先进先出
//&栈用数组表示，分别进行进队和出队操作
type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

//*出队
func (this *CQueue) DeleteHead() int {
	//*outStack存在元素，从尾部获取
	if len(this.outStack) != 0 {
		v := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		return v
	}

	//*inStack元素数量为空
	if len(this.inStack) == 0 {
		return -1
	}

	//?逆序复制inStack的值到outStack
	for i := len(this.inStack) - 1; i >= 0; i-- {
		this.outStack = append(this.outStack, this.inStack[i])
	}

	v := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return v
}
