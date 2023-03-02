package main

import (
	"fmt"
)

//*统计一个数字在排序数组出现的次数
//?一般思路：遍历，变量count记录目标数字的次数；时间复杂度O(N)

func main() {
	nums := []int{1, 2, 4, 5, 6}
	fmt.Println(binarySearch(nums, 2))
}

//~数组是有序的，可以考虑使用二分查找
//&二分查找实现-线性
func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums)

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

/*
*使用二分查找分别确定目标数字的左右边界，出现次数计算即r-l-1
 */
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	//*确定右边界 right=l
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	right := l
	if r >= 0 && nums[r] != target { //*目标数字不存在
		return 0
	}
	//*确定作边界left=r
	l = 0
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	left := r

	return right - left - 1
}

//!优化
func BestSeacher(nums []int, target int) int {
	seacher := func(target int) int {
		l := 0
		r := len(nums) - 1

		for l <= r {
			mid := (l + r) / 2
			if nums[mid] <= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		return l
	}

	return seacher(target) - seacher(target-1)
}

//*查找递增数字中缺失数字，数字长度n-1，每个数字0~n-1范围
//&使用二分查找
func missingNumber(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}
