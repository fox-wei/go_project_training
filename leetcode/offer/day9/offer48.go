package day9

//*最长不含重复字符串的子字符串
func lengthOfLongestSubstring(s string) int {
	dic := make(map[string]int) //*记录字符最后出现的位置
	res, tmp := 0, 0

	for j := 0; j < len(s); j++ {
		ss := string(s[j])
		i, ok := dic[ss] //*获取s[j]索引
		if !ok {
			i = -1
		}
		dic[ss] = j //*更新map
		if tmp < j-i {
			tmp++
		} else {
			tmp = j - i
		}

		if res < tmp {
			res = tmp
		}
	}

	return res
}
