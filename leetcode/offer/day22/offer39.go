package day22

//*数组中出现次数超过一半的数字（该数字一定存在）
//*1.哈希表:数字-统计次数；次数大于长度一般返回；
//*2.排序：中间值一定是目标数字（数字次数超过长度的一半）
/*
 * 摩尔投票法：
 *"众数”标记votes+1；非“众数”标记votes-1；前a(a<n)个之和==0(votes==0)；
 *target=s[a+1]；最终结果一定votes>0；返回target。
 */
func majorityElement(nums []int) int {
	target := nums[0] //*目标数
	votes := 1        //*投票
	for i := 1; i < len(nums); i++ {
		if votes == 0 {
			target = nums[i]
		}
		if nums[i] == target {
			votes += 1
		} else {
			votes -= 1
		}

	}
	return target
}
