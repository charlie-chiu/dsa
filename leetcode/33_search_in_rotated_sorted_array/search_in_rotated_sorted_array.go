package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r-l)/2

		if target < nums[0] && nums[0] < nums[m] {
			l = m + 1 // M inside +INF
		} else if target >= nums[0] && nums[0] > nums[m] {
			r = m // M inside -INF
		} else if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m
		} else {
			return m
		}
	}

	return -1
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
