package main

import (
	"math"
)

func SnowballNumberArithmetic(i int) [][]int {

	var res [][]int

	max := int(math.Ceil(float64(i) / 2))
	for k := 1; k < max; k++ {

		b, c := getBAndCFromQuadraticEquation(k, i)
		x := quadraticEquationPositif(1, float64(b), float64(c))
		_, decimal := math.Modf(x) // check jika x adalah float tanpa decimal
		if decimal == 0 {
			// create list to x number
			list := []int{k}
			for j := k + 1; j <= int(x)+k-1; j++ {
				list = append(list, j)
			}

			res = append(res, list)
		}
	}

	res = append(res, []int{i})
	return res
}

func getBAndCFromQuadraticEquation(a, i int) (int, int) {
	// persamaan kuadat ax2+bx+c
	// untuk mengambil nilai b dan c dari persamaan kuadrat
	// nilai a selalu 1 untuk deret aritmatik dengan b = 1

	return (2 * a) - 1, (2 * i) * -1
}

func quadraticEquationPositif(a, b, c float64) float64 {

	D := (b * b) - (4 * a * c)
	var x1 float64

	if D > 0 {
		x1 = (-b + math.Sqrt(D)) / (2 * a)

	} else if D == 0 {
		x1 = (-b + math.Sqrt(D)) / (2 * a)

	} else {
		x1 = -b / (2 * a)
	}

	return x1
}

