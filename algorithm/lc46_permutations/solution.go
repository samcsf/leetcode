package lc46

// Permute 全排列
func Permute(nums []int) [][]int {
	var (
		tmp []int
		res [][]int
	)
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
	for i, n := range nums {
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
