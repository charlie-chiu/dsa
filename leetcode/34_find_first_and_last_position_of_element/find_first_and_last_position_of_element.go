package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	lower := lowerBound(nums, target)
	upper := upperBound(nums, target) - 1
	if lower >= len(nums) || nums[lower] != target {
		return []int{-1, -1}
	}

	return []int{lower, upper}
}

func lowerBound(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r-l)/2

		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func upperBound(nums []int, target int) int {
	l := 0
	r := len(nums)

	for l < r {
		m := l + (r-l)/2
		//binarySearchVisualization(nums, l, r, m)

		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func binarySearchVisualization(elements []int, l, r, m int) {
	fmt.Printf("L: %d M: %d R: %d", l, m, r)
	fmt.Println()
	for _, element := range elements {
		fmt.Printf(" %2d", element)
	}
	fmt.Println()
	for i := range elements {
		switch i {
		case l, r, m:
			fmt.Printf("%s", "  ^")
		default:
			fmt.Print("   ")

		}
	}
	fmt.Println()
	for i := range elements {
		switch i {
		case l:
			fmt.Print("  L")
		case r:
			fmt.Print("  R")
		case m:
			fmt.Print("  M")
		default:
			fmt.Print("   ")
		}
	}
	fmt.Println()
	fmt.Println()
}
