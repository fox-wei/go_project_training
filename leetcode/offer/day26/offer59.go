package day26

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	if nums == nil || k <= 0 {
		return res
	}

	window := make([]int, 0, k) //*双端队列，最大元素在头
	for i, num := range nums {
		for len(window) > 0 && num > nums[window[len(window)-1]] {
			//*新加入的值大于双向队列的最后一个数，最后元素出队
			window = window[:len(window)-1]
		}
		window = append(window, i)

		if i >= k-1 { //*形成窗口
			res = append(res, nums[window[0]])
			if window[0] <= i-k+1 {
				//*如果最大值所在下标小于等于i-k+1，出队
				window = window[1:]
			}
		}
	}
	return res
}

//?队列的最大值
type MaxQueue struct {
	queue    []int //*普通队列
	maxQueue []int //*维护最大值在队头
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	//*最大值队列维护
	//*保持队列是递减状态
	//*value > 队尾元素，队尾元素出队
	for len(this.maxQueue) > 0 && value > this.maxQueue[len(this.maxQueue)-1] {
		this.maxQueue = this.maxQueue[:len(this.maxQueue)-1]
	}
	this.maxQueue = append(this.maxQueue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:] //*出队
	if res == this.maxQueue[0] {
		this.maxQueue = this.maxQueue[1:]
	}
	return res
}
