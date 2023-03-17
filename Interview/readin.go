package main

func minSubarray(nums []int, p int) int {
	length := len(nums)
	    presum := make([]int, length + 1)
	for i, v := range nums {
	    presum[i + 1] = (presum[i] + v) % p
	}
	if presum[length] == 0 {
	    return 0
	}
    
	res := -1
	for l := 1; l < length; l++ {
	    for start := 0; start + l <= length; start++ {
		remain := (presum[start + l] - presum[start]) % p
		if remain == presum[length] {
		    return l
		}
	    }
	}
	    return res
    }
     

func main() {
	minSubarray([]int{3,1,4,2}, 6)
}