package main

//*字符串替换
func replaceSpace(s string) string {
	var res []rune

	for _, v := range s {
		if v == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, v)
		}
	}

	s = s[:]

	return string(res)
}
