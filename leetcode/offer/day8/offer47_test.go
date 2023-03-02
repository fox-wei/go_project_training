package day8

import (
	"fmt"
	"testing"
)

func Test_maxValue(t *testing.T) {
	s := [][]int{{1, 3, 1}, {1, 5, 1}}
	// m := len(s[0])
	// for i:=1; i<m;
	res := maxValue(s)
	fmt.Println(s)
	fmt.Println(res)
}
