package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	var m int
	for l < r {
		m = l + (r-l)/2

		if nums[m] == target {
			return m
		} else if target > nums[m] {
			l = m + 1
		} else if nums[m] > target {
			r = m
		}
	}

	return r
}
