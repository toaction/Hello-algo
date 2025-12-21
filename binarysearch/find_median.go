package binarysearch

/*
4. 寻找两个正序数组的中位数：https://leetcode.cn/problems/median-of-two-sorted-arrays/

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出并返回这两个正序数组的 中位数 。
*/

// 在两个有序数组中找第 k 小的数
// 每次从两个数组中的一个淘汰 k/2 个数，更新 k 的值为 k-k/2，直到 k=1
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := (m + n) / 2
	if (m+n)%2 == 1 {
		return getMinK(nums1, nums2, k+1)
	}
	return (getMinK(nums1, nums2, k) + getMinK(nums1, nums2, k+1)) / 2
}

func getMinK(nums1 []int, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	// 两个数组的左边界
	first, second := 0, 0
	for first < m && second < n && k > 1 {
		mid1 := first + k/2 - 1
		// 如果可用数组长度小于 k/2，将其指向最后一个元素
		if mid1 >= m {
			mid1 = m - 1
		}
		mid2 := second + k/2 - 1
		if mid2 >= n {
			mid2 = n - 1
		}
		// 第一个可用数组的前半部分元素都小于等于第 k 小元素
		if nums1[mid1] <= nums2[mid2] {
			k -= mid1 - first + 1
			first = mid1 + 1
		} else {
			k -= mid2 - second + 1
			second = mid2 + 1
		}
	}
	ans := 0
	// 第 k 小的数在第二个有序数组中
	if first == m {
		ans = nums2[second+k-1]
	} else if second == n {
		ans = nums1[first+k-1]
	} else {
		ans = min(nums1[first], nums2[second])
	}
	return float64(ans)
}
