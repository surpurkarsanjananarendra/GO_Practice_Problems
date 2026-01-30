package main

import "fmt"

func Incrementer(n []int) []int {
	for i := range n {
		n[i] = (n[i] + (i + 1)) % 10
	}
	return n
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Incrementer(nums))
}
