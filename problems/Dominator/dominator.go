package main

import "fmt"

func Dominator(a []int) int {
	half := int(len(a) / 2)
	res := make(map[int]int)
	for i := range a {
		res[a[i]] += 1
	}
	for k, v := range res {
		if v >= half {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(Dominator([]int{3, 4, 3, 2, 3, 1, 3, 3}))
}
