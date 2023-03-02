package day12

//*实现数组奇数在前半部分，偶数在后半部分
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		for i < j && !isEven(nums[i]) {
			i++
		}
		for i < j && isEven(nums[j]) {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func isEven(n int) bool {
	return n%2 == 0
}
