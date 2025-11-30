package sort

// 堆排序
//  1. 输入数组并建立大顶堆
//  2. 将堆顶元素与堆底元素交换，交换后堆大小减一；然后调整堆结构
//  3. 重复交换和堆调整，即可完成数组排序
func HeapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		siftDown(arr, i, 0)
	}
}

// 自上到下调整堆
func siftDown(arr []int, n int, idx int) {
	for idx < n/2 {
		tmp := idx
		left, right := 2*idx+1, 2*idx+2
		if left < n && arr[left] > arr[idx] {
			idx = left
		}
		if right < n && arr[right] > arr[idx] {
			idx = right
		}
		if tmp == idx {
			break
		}
		arr[tmp], arr[idx] = arr[idx], arr[tmp]
	}
}
