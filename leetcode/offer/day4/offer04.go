package main

//?n*m二维数组查找；数组行列都是递增；判断元素是否在数组中
//*二维数组反转45°看，15为例，是第一行的最大值，对应列的最小值；
//&解题思路：从右上角开始：f>t，说明可能在所在行j--；f<t，不在当前行，i++
/*
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
*/

//?从“右上角”开始
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix[0]) - 1

	for j > 0 && i < len(matrix) {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}

	return false

}
