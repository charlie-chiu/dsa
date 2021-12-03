package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	var index1 = 0
	var index2 = len(numbers) - 1

	//The tests are generated such that there is exactly one solution.
	for {
		switch sum := numbers[index1] + numbers[index2]; {
		case sum < target:
			index1++
		case sum > target:
			index2--
		case sum == target:
			return []int{index1 + 1, index2 + 1}
		}
	}
}
