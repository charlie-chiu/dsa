package main

func singleNonDuplicate(nums []int) int {
	// 出現單值之前，奇位偶位相同，出現後則不相同
	l := 0
	r := len(nums)
	var m int

	for l < r {
		m = l + (r-l)/2

		// target found
		if m == 0 || m == len(nums)-1 || (nums[m-1] < nums[m] && nums[m] < nums[m+1]) {
			return nums[m]
		}

		if m%2 == 0 {
			m++
		}

		if nums[m] == nums[m-1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return 0
}
