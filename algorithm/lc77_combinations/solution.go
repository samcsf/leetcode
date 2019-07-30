package lc77

// Combine 组合
func Combine(n int, k int) [][]int {
	var tmp []int
	var nums []int
	var res [][]int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	getPair(k, nums, &tmp, &res)
	return res
}

func getPair(k int, nums []int, tmp *[]int, res *[][]int) {
	if len(*tmp) >= k {
		t := make([]int, len(*tmp))
		copy(t, *tmp)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		tt := append(*tmp, nums[i])
		getPair(k, nums[i+1:], &tt, res)
	}
}
