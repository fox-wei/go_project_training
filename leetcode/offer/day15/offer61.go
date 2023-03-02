package day15

import "sort"

//*扑克牌中的顺子：判断5张牌是否连续
//*特殊点：大小王是任意牌（0表示）
//*构成顺子的条件：无重复数字（大小王除外），最大值-最小值<5

//!方法1：集合+遍历
func isStraight(nums []int) bool {
	repeat := make(map[int]bool) //集合
	mi, ma := 14, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] > ma {
			ma = nums[i]
		}

		if nums[i] < mi {
			mi = nums[i]
		}

		if repeat[nums[i]] {
			return false
		}

		repeat[nums[i]] = true
	}

	return ma-mi < 5
}

//!方法2：排序+遍历
//&排序后，遍历计算大小王数量n；判断num[4]-num[n]<5
func sStraight2(nums []int) bool {
	sort.Ints(nums) //*升序
	joker := 0      //*记录小大小王数量
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker+1] < 5
}
