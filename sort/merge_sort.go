package sort

// 归并排序：基于分治思想，不断将数组从中点分开，将长数组的排序问题转为短数组的排序问题
// 当子数组长度为 1 时终止划分，开始合并，持续地将左右两个较短的有序数组合并为一个较长的有序数组
func MergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := (right-left)/2 + left
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// 合并两个有序数组
func merge(arr []int, left int, mid int, right int) {
	tmp := make([]int, right-left+1)
	k := 0
	l, r := left, mid+1
	for l <= mid && r <= right {
		if arr[l] < arr[r] {
			tmp[k] = arr[l]
			l++
		} else {
			tmp[k] = arr[r]
			r++
		}
		k++
	}
	for l <= mid {
		tmp[k] = arr[l]
		l++
		k++
	}
	for r <= right {
		tmp[k] = arr[r]
		r++
		k++
	}
	for i := left; i <= right; i++ {
		arr[i] = tmp[i-left]
	}
}
