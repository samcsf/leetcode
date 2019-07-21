package lc40

import (
	"sort"
)

// CombinationSum II 组合总和 II 不允许重复使用
// 基于lc39组合总和1的解法，增加数组推移（下组候选数不包括重复）
// 完成一组计算后推移指针，跳开重复前缀递归
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var tmpRes []int
	sort.Ints(candidates)
	cb(candidates, target, &tmpRes, &res)
	return res
}

func cb(nums []int, target int, tmpRes *[]int, finalRes *[][]int) {
	for i := 0; i < len(nums); i++ {
		// 防止越界
		for i > 0 && i < len(nums)-1 && nums[i] == nums[i-1] {
			i++
		}
		n := nums[i]
		if n >= target {
			if n == target {
				t := make([]int, len(*tmpRes))
				copy(t, *tmpRes)
				t = append(t, n)
				// 注意一定要复制新的slice，不然会被后面改到
				*finalRes = append(*finalRes, t)
			}
			return
		}
		if i+1 < len(nums) {
			tr := append(*tmpRes, n)
			cb(nums[i+1:], target-n, &tr, finalRes)
		}
	}
}
