package lc26

// RemoveDuplicates remove duplicate elements
// 使用快慢两个指针：
// 1. 慢指针用于编织新的非重数组
// 2. 快指针用于探路
// 当遇到重复元素跳过，然后将元素复制到慢指针后面
// 实现一次循环将有序数组去重
// #双指针 #快慢指针
func RemoveDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
