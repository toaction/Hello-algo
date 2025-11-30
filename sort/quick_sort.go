package sort

// 快速排序：基于分治思想
// 每轮选择一个哨兵，将所有小于哨兵的元素移动到其左侧，大于哨兵的元素移动到其右侧
func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := partition(arr, left, right)
	QuickSort(arr, left, mid-1)
	QuickSort(arr, mid+1, right)
}

// 哨兵划分
func partition(arr []int, left int, right int) int {
	// 以 arr[left] 作为哨兵
	i, j := left, right
	for i < j {
		// 最后要将 arr[i] 与哨兵交换，arr[i] 最后的值一定要 <= 哨兵，因此先从右向左遍历找到小于哨兵的元素
		for i < j && arr[j] >= arr[left] {
			j--
		}
		for i < j && arr[i] <= arr[left] {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[i], arr[left] = arr[left], arr[i]
	return i
}
