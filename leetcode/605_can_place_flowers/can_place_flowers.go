package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		return flowerbed[0] == 0 && n == 1 || n == 0
	}

	//strategy : 只要左右兩側沒花就種下去
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		switch i {
		case 0:
			if flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n -= 1
			}
		case len(flowerbed) - 1:
			// last plot
			if flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n -= 1
			}
		default:
			if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n -= 1
			}
		}
	}

	return n == 0
}
