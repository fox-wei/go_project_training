package day12

//*求递增数组和为s的两个数字
//*使用双指针
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		s := nums[i] + nums[j]
		if s < target {
			i++
		} else if s > target {
			j--
		} else {
			return []int{nums[i], nums[j]}
		}
	}

	return nil
}
