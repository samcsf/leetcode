package lc47

import "sort"

// PermuteUnique 全排列II
// 注意每步循环前跳过相同的元素
func PermuteUnique(nums []int) [][]int {
	var (
		tmp []int
		res [][]int
	)
	sort.Ints(nums)
	do(nums, &tmp, &res)
	return res
}

func do(nums []int, tmp *[]int, res *[][]int) {
	if len(nums) == 1 {
		t := make([]int, len(*tmp))
		copy(t, *tmp)
		t = append(t, nums[0])
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		for i < len(nums)-1 && nums[i+1] == nums[i] {
			i++
		}
		n := nums[i]
		tt := append(*tmp, n)
		do(popIndex(nums, i), &tt, res)
	}
}

func popIndex(nums []int, i int) []int {
	var trail []int
	trail = append(trail, nums[:i]...)
	if i < len(nums)-1 {
		trail = append(trail, nums[i+1:]...)
	}
	return trail
}
