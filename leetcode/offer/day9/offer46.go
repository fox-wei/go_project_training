package day9

func translateNum(num int) int {

	s := string(rune(num)) //*转化成字符串

	a, b := 1, 1 //*dp状态表

	for i := 2; i < len(s); i++ {
		tmp := s[i-2 : i]
		c := 0
		if tmp >= "10" && tmp <= "25" {
			c = a + b
		} else {
			c = a
		}
		b = a
		a = c

	}
	return a
}
