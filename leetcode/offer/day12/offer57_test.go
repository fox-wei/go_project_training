package day12

import (
	"fmt"
	"strings"
	"testing"
)

func Test_twoSum(t *testing.T) {
	s := " I am a  student. "
	s = strings.Trim(s, " ")
	fmt.Println(s)
	ss := ""
	i, j := len(s)-1, len(s)

	for i >= 0 {
		for i >= 0 && string(s[i]) != " " {
			i--
		}
		fmt.Println(i, string(s[i+1:j]))
		ss += string(s[i+1 : j])
		if i <= 0 {
			break
		}

		for string(s[i]) == " " {
			i--
		}
		ss += " "
		j = i + 1

	}
	fmt.Println(ss)
}
