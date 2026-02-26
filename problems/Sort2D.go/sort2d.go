package main

import "fmt"

func Snail(snaipMap [][]int) []int {
	var res []int
	for _, nums := range snaipMap {
		for _, n := range nums {
			res = append(res, n)
		}
	}
	for i := 0; i < len(res)-1; i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(Snail([][]int{{1, 2, 3},
		{6, 7, 8},
		{5, 4, 9}}))
}
