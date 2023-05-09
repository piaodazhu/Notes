package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)

	res := math.MaxInt32
	minval := min(nums[0], nums[n - k] >> 1)
	maxval := max(nums[n - 1 - k], nums[n - 1] >> 1)
	res = min(res, maxval - minval)

	minval = min(nums[0] << 1, nums[k])
	maxval = max(nums[k - 1] << 1, nums[n - 1])
	res = min(res, maxval - minval)
	for mul2 := 1; mul2 < k; mul2++ {
		minval = min(min(nums[0] << 1, nums[mul2]), nums[n - k + mul2] >> 1)
		maxval = max(max(nums[mul2 - 1] << 1, nums[n - k + mul2 - 1]), nums[n - 1] >> 1)
		res = min(res, maxval - minval)
	}
	fmt.Println(res)
}
func min(x, y int) int {
	if x < y {
		return x 
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x 
	}
	return y
}