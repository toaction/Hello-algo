package sort

// 快速排序：基于分治思想
// 每轮选择一个哨兵，将所有小于哨兵的元素移动到其左侧，大于哨兵的元素移动到其右侧
func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	piv := partition(arr, left, right)
	QuickSort(arr, left, piv)
	QuickSort(arr, piv+1, right)
}

// Hoare 划分
func partition(arr []int, left int, right int) int {
	// 以 arr[left] 作为哨兵
	pivot := arr[left]
	i, j := left-1, right+1
	for i < j {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 左区间 <= pivot，右区间 >= pivot，但划分点（左区间的右端点）不一定在最终位置
	return j
}
