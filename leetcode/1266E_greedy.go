package main

// TODO

func minTimeToVisitAllPoints(points [][]int) int {
	total := 0
	for i := 1; i < len(points); i++ {
		// dx = x2 - x1
		dx := abs(points[i][0] - points[i-1][0])

		// dy = y2 - y1
		dy := abs(points[i][1] - points[i-1][1])

		// cover most distance
		total += max(dx, dy)
	}
	return total
}
