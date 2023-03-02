package day9

import (
	"fmt"
	"testing"
)

func Test_translateNum(t *testing.T) {
	num := 1234
	s := fmt.Sprint(num)
	fmt.Println(len(s), s)
	s2 := s[0:2]
	s1 := s[2:4]
	fmt.Println(s2 < s1)
}
