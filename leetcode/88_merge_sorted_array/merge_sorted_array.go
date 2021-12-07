package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		return
	}

	p1 := m - 1
	p2 := n - 1
	insertAt := m + n - 1

	for insertAt >= 0 {
		if p1 < 0 {
			nums1[insertAt] = nums2[p2]
			p2--
		} else if p2 < 0 {
			nums1[insertAt] = nums1[p1]
			p1--
		} else {
			if nums2[p2] > nums1[p1] {
				nums1[insertAt] = nums2[p2]
				p2--
			} else {
				nums1[insertAt] = nums1[p1]
				p1--
			}
		}

		insertAt--
	}

	return
}
