package main

import "fmt"

func ReverseSeq(n int) []int {
	res := make([]int, n)
	c := n
	for i := 0; i < n; i++ {
		res[i] = c
		c--
	}
	return res
}

func main() {
	fmt.Println(ReverseSeq(5))
}
