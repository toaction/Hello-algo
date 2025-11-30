package sort

// 选择排序：每轮从未排序的元素中选择最小的元素，将其放在已排序的元素末尾
func SelectSort(arr []int) {
	// 未排序区间 [i, n-1]
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// 将区间最小元素放在已排序的元素后面，即区间最左端
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
