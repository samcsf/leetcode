package lc78

// Subsets 子集
func Subsets(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	for i := 0; i < 1<<uint(l); i++ {
		var tmp []int
		for j := 0; j < l; j++ {
			sel := (i >> uint(l-j-1)) & 1
			if sel == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		res = append(res, tmp)
	}
	return res
}
