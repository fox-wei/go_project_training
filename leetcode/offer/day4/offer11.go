package main

//*寻找翻转数组的最小值
//?遍历实现
func minArray(numbers []int) int {

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] { //反转点两边都是升序
			return numbers[i]
		}
	}

	return numbers[0]
}

//!使用二分查找
//?翻转数组的特点是边界点两边左边界大于右边界
//&使用二分查找mid>l；边界点在右边界l=mid+1；mid<r,在左边界,mid=right；
//!本题难点是元素重复；如果mid=right，是无法进行边界划分；所以要rihgt-1
func minArrayByBinary(numbers []int) int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		mid := (l + r) / 2

		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}

	return numbers[l]
}
