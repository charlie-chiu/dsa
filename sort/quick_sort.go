package sort

import "log"

const devMsg = "%10s --- seq:%v markerL:%d markerR:%d"

// implement app version
// pivot at rightmost
func quickSort(input []int) {
	log.SetFlags(log.Lshortfile)
	if len(input) <= 1 {
		return
	}

	lastIndex := len(input) - 1
	var pivot = lastIndex
	var l = 0
	var r = lastIndex
	var sortedIndex int
	//log.Printf(devMsg, "START", input, l, r)

	for l < r {
		// left marker find value ge than pivot
		for !(input[l] >= input[pivot]) {
			l++
		}

		// right marker find value less than pivot
		for !(input[r] < input[pivot]) && r > 0 {
			r--
		}

		if l < r {
			input[l], input[r] = input[r], input[l]
		}
		//log.Printf(devMsg, "AFTER LOOP", input, l, r)
	}

	input[l], input[pivot] = input[pivot], input[l]
	sortedIndex = l

	//log.Printf(devMsg, "FINAL", input, l, r)
	//log.Printf("sortedValue: %d at index %d ", input[sortedIndex], sortedIndex)
	//log.Printf("")
	leftside := input[0:sortedIndex]
	rightside := input[sortedIndex+1:]
	quickSort(leftside)
	quickSort(rightside)
}
