package lc754

import "math"

// ReachNumber 到达数字
// 这题没什么很通用的方法论，就是看数字规律
// 首先找到最接近 target 的步数 n, (比如 1+2+3 最接近 5)
// 求最接近的 delta = 1 + ... + n - target
//   * 关键结论 delta 如果为偶数，那么必然可以通过修改 1+...+n 中的符号得到解
// 所以，问题转化为使 delta 变为偶数，当 delta 为奇数， 开始尝试 n + 1
// 如果 n + 1 为奇数 delta 为奇数， 奇 + 奇 = 偶，那么 n + 1 为解
// 如果 n + 1 为偶数 delta 为奇数， 奇 + 偶 = 奇，那么将回到上面的问题
// 继续往下走再令 n + 2 即可(偶数加 1 必然奇数) 奇 + 偶 + 奇 = 偶， 那么 n + 2为解
func ReachNumber(target int) int {
	target = int(math.Abs(float64(target)))
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 != 0 {
		return k + 1 + k%2 // k%2 快速判断奇偶， +0 / +1
	}
	return k
}
