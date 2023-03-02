package main

import "fmt"

func main() {
	s := make(map[byte]bool)

	fmt.Println(s['a'])
}

//?指出字符串第一个只出现一次的字符；建立字符-次数关系。
func firstUniqChar(s string) byte {
	var m [26]int

	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}
