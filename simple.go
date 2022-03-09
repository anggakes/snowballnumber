package main

func SnowballNumberSimple(N int) (R [][]int) {

	for k := 1; k <= N; k++ {

		sum := 0
		var list []int
		for j := k; j <= N; j++ {

			sum = sum + j
			list = append(list, j)
			if sum >= N {
				break
			}
		}
		if sum == N {
			R = append(R, list)
		}
	}
	return R
}
