package day20

//*求整数（二进制表示）中的1的个数：要求时间复杂度O(N)，空间O(1)
//*使用与运算：与0为零；1与1为1；然后右移（去掉最后一位）
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1 //?右移即去掉最后一位
	}
	return count
}
