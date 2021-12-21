package leetcode

import "fmt"

func BinarySearchVisualization(elements []int, l, r, m int) {
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
