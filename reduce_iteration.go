package main

import (
	"math"
)

func SnowballNumberReduceIteration(N int) [][]int {
	var res [][]int

	max := int(math.Ceil(float64(N) / 2))

	currSum := 1
	currList := []int{1}
	currJ := 2

	for k := 1; k <= max; k++ {

		for j := currJ; j <= max; j++ {
			currJ = j
			currSum += currJ // gunakan sum sebelumnya untuk mengurangi jumlah iterasi

			if currSum <= N {
				if currJ > currList[len(currList)-1]{ // append list jika belum ada pada currList
					currList = append(currList, j)
				}
				if currSum == N {
					res = append(res, currList)
					break
				}
			} else if currSum > N{
				break
			}
		}

		// hapus list dan kurangi sum dengan nilai k sekarang
		currSum -= k
		currSum -= currJ
		currList = currList[1:]

	}

	res = append(res, []int{N})

	return res
}
