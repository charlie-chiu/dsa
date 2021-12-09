package main

func mySqrt(x int) int {
	if x == 0 {
		return x
	}

	l := 1
	r := x

	for l <= r {
		mid := l + (r-l)/2
		square := mid * mid
		if square == x {
			return mid
		} else if square > x {
			r = mid - 1
		} else if square < x {
			l = mid + 1
		}
	}

	return r
}

//func mySqrt(x int) int {
//	if x == 0 {
//		return 0
//	}
//
//	l := 1
//	r := x
//
//	for l < r {
//		//mid := (l+r)/2
//		mid := l + (r-l)/2+1
//		//fmt.Printf("X: %3d |L: %3d |R: %3d |MID: %3d |MID^2: %3d\n",x, l, r, mid, mid * mid)
//		fmt.Printf("X: %3d |%3d---%3d---%3d |MID^2: %3d\n",x, l, mid, r, mid * mid)
//		if mid * mid > x {
//			r = mid-1
//		} else {
//			l = mid
//		}
//	}
//
//	return l
//}
