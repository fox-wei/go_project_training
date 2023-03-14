package offer24

import (
	"fmt"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	s := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := spiralOrder(s)
	fmt.Println(res)
}
