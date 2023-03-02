package day12

import "strings"

//*反转单词顺序：输入一个英文句子，反转单词顺序但是单词顺序不变
//&使用双指针i,j,从末尾开始遍历
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	ss := ""
	i, j := len(s)-1, len(s)

	for i >= 0 {
		for i >= 0 && string(s[i]) != " " { //*确定单词
			i--
		}

		ss += s[i+1 : j]

		if i <= 0 {
			break
		}

		for string(s[i]) == " " {
			i--
		}

		j = i + 1
		ss += " "
	}

	return ss
}
