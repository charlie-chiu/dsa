package main

import "sort"

func findMinArrowShots(points [][]int) int {
	// 策略，每次選一能射爆最多氣球的位置射箭，直到沒有氣球為止
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	shots := 0

	for len(points) > 0 {
		shotAt := points[0][1]

		var afterShot [][]int
		for _, point := range points {
			if point[0] > shotAt || shotAt > point[1] {
				afterShot = append(afterShot, point)
			}
		}
		points = afterShot

		shots++
	}

	return shots
}
