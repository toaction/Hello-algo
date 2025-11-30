package sort

// 插入排序：在未排序区间选择一个基准元素，将其插入到已排序区间的正确位置
func InsertionSort(arr []int) {
	n := len(arr)
	// 已排序区间为 [0, i-1]
	for i := 1; i < n; i++ {
		base := arr[i]
		j := i - 1
		// 倒序遍历，找到要插入的位置
		for j >= 0 && arr[j] > base {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = base
	}
}
