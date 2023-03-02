package day8

//*礼物的最大值，棋盘n*m每个格子都有价值，从左上角每次只能向右或向下移动；移动到右下角结束
//*使用动态规划，先对第一行和第一列进行更新，然后开始循环，修改礼物价值；得到最终结果

func maxValue(grid [][]int) int {
	n := len(grid)    //*行
	m := len(grid[0]) //?列
	//?更新第一行的值
	for i := 1; i < m; i++ {
		grid[0][i] += grid[0][i-1]
	}
	//*更新第一列
	for i := 1; i < n; i++ {
		grid[i][0] += grid[i-1][0]
	}

	//?更新棋盘的值
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i-1][j] > grid[i][j-1] { //*上大于左边
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[n-1][m-1]
}
