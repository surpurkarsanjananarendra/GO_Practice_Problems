package main

import "fmt"

func main() {
	/*
	   c := make(chan string)
	   // var c chan string
	   go subRountine(c)
	   fmt.Println(<-c)
	*/
	var sli []int
	fmt.Println(sli)
	slic := make([]int, 5)
	fmt.Println(slic)
	slic[0] = 1
	slic[2] = 3
	fmt.Println(slic)
}

func subRountine(c chan string) {
	// c <- "hello"
}
