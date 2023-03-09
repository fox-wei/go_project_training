package day20

//?加法运算使用位运算符实现
//*使用异或运算符
//!难点：进位问题
func add(a, b int) int {
	c := 0
	for b > 0 {
		c = (a & b) << 1 //*进位计算
		a ^= b           //*异或运算
		b = c            //?更新进位

	}
	return a
}
