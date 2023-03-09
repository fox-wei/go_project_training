package day21

import "fmt"

//*I：一个数组只有两个数字(x,y)仅出现一次；其他数字出现两次
//?要求时间O(N)，空间是O(1)
//*使用异或运算（满足交换律）：同为0；异为1
//^异或运算和分组：最终结果是x^y；x^y的结果必定有1;获取第一个低位1根据这个1的位置进行分组
//*假设数组仅一个数字仅出现一次，其他两次
func sigleOnly(nums []int) int {
	res := 0
	for _, k := range nums {
		res ^= k
	}
	return res
}
func singleNumbers(nums []int) []int {
	x, y, n, mask := 0, 0, 0, 1
	//*计算x^y
	for _, v := range nums {
		n ^= v //&x^y
	}
	fmt.Println(n)
	//*获取x^y的低位第一个1
	for n&mask == 0 {
		mask <<= 1
	}
	fmt.Println(mask)

	//*mask得到后进行分组
	for _, v := range nums {
		if v&mask == 0 { //*必须是0
			x ^= v
		} else {
			y ^= v
		}
		fmt.Println(v&mask, x, y)
	}
	return []int{x, y}
}

//!II:数组只有一个数字出现一次，其他数字出现三次
