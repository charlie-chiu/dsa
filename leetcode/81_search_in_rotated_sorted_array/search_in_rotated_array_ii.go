package main

func search(nums []int, target int) bool {
	l := 0
	// 需要用到右端點的值，故使用左閉右閉的寫洲
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2

		//leetcode.BinarySearchVisualization(nums, l, r, m)

		if nums[m] == target {
			return true
		}

		if nums[l] == nums[m] {
			//fmt.Println("無法判斷哪一區間增序")
			// 無法判斷哪一區間增序
			l++
		} else if nums[m] <= nums[r] {
			//fmt.Println("右區間增序")
			// 右區間增序
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m
			}
		} else {
			//fmt.Println("左區間增序")
			// 左區間增序
			if target >= nums[l] && target < nums[m] {
				r = m
			} else {
				l = m + 1
			}
		}
	}

	return false
}
