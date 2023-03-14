package offer24

//*顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}

	row, col := len(matrix), len(matrix[0]) //*行，列
	left, right, up, down := 0, col-1, 0, row-1
	//*从左到右，从上到下，从右到左，从下到上重复调整四个边界
	for left <= right && up <= down {
		for i := left; i <= right; i++ { //?左到右
			res = append(res, matrix[up][i])
		}

		for i := up + 1; i <= down; i++ { //*上到下

			res = append(res, matrix[i][right])
		}

		if up < down { //*右到左
			for i := right - 1; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
		}

		if left < right { //*下到上
			for i := down - 1; i > up; i-- {
				res = append(res, matrix[i][left])
			}
		}

		left++
		right--
		up++
		down--
	}
	return res
}
