package sort

// 冒泡排序：依次比较相邻元素大小，如果“左元素 > 右元素”就交换二者。
// 遍历完成后，最大的元素会被移动到数组的最右端。
func BubbleSort(arr []int) {
	n := len(arr)
	// 未排序区间为 [0, i]
	for i := n - 1; i > 0; i-- {
		// 将区间最大值移动到区间最右端
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
