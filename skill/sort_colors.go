package skill

/*
75. 颜色分类: https://leetcode.cn/problems/sort-colors/

给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，
原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

示例 1：
	输入：nums = [2,0,2,1,1,0]
	输出：[0,0,1,1,2,2]
*/

// 双指针，先将整数2交换到尾部，然后将整数0交换到前面
// 整数2交换后，值可能还等于2，继续交换
func SortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
