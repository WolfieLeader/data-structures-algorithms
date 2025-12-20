package main

// I: gain = [-5,1,5,0,-7]       => Ex: alts = [0,-5,-4,1,1,-6],          high->1 => O: 1
// I: gain = [-4,-3,-2,-1,4,3,2] => Ex: alts = [0,-4,-7,-9,-10,-6,-3,-1], high->1 => O: 0

func largestAltitude(gain []int) int {
	sum, highest := 0, 0
	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		highest = max(highest, sum)
	}
	return highest
}
