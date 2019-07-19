package lc39

import (
	"sort"
)

// CombinationSum 组合总和
// 分析路径，利用回溯对所有可能的路径进行递归，遇到匹配值回调
// 事先进行数组排序方便检查到大于总和的情况后放弃后续的递归
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var tmpRes []int
	sort.Ints(candidates)
	cb(candidates, target, &tmpRes, &res)
	return res
}

func cb(nums []int, target int, tmpRes *[]int, finalRes *[][]int) {
	for i := 0; i < len(nums); i++ {
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
		tr := append(*tmpRes, n)
		cb(nums[i:], target-n, &tr, finalRes)
	}
}
