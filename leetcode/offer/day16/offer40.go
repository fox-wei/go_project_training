package day16

//*返回数组的k最小值
//*数组排序实现-快速排序实现
func getLeastNumbers(arr []int, k int) []int {
	//*实现快速排序
	var quickSort func(arr []int, l, r int)

	quickSort = func(arr []int, l, r int) {
		if l >= r {
			return
		}

		i, j := l, r
		for i < j {
			for i < j && arr[j] >= arr[l] {
				j--
			}

			for i < j && arr[i] <= arr[j] {
				i++
			}

			arr[i], arr[j] = arr[j], arr[i]
		}

		arr[l], arr[i] = arr[i], arr[l]
		quickSort(arr, l, i-1)
		quickSort(arr, i+1, r)
	}
	quickSort(arr, 0, len(arr)-1)
	return arr[:k]
}
