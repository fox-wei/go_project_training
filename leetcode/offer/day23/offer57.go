package day23

//*输入一个整数；输出所有和为target的 连续正整数序列（至少两个数）；

//!语法错误
func findContinuousSequence1(target int) [][]int {
	//*关键词：连续正整数序列
	res := [][]int{}
	i, j, s := 1, 2, 3 //*使用双指针,s表示序列和

	for i < j {

		if s < target {
			//*j向右移动，并加上新数字
			j++
			s += j
		} else if s > target {
			//*减去第一数字，i向右
			s -= i
			i++
		} else {
			//*找到目标
			tmp := make([]int, j-i+1)
			for n := i; n <= j; n++ {
				tmp[n-i] = n
			}
			res = append(res, tmp) //!发生内存泄漏，忘记更新s
			s -= i
			i++
		}

	}
	return res
}

func findContinuousSequence(target int) [][]int {
	res := [][]int{}
	i, j, s := 1, 2, 3
	for i < j {
		if target > s {
			j++
			s += j
		} else {
			if s == target {
				tmp := make([]int, j-i+1)
				for m := i; m <= j; m++ {
					tmp[m-i] = m
				}
				res = append(res, tmp)
			}
			s -= i
			i++
		}
	}

	return res
}
