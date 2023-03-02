package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
		if m[v] >= 2 {
			return v
		}
	}

	return -1
}
