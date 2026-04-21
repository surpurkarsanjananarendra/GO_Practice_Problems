package main

import "fmt"

func subRoutine(num int, c chan int) {
	x, y := 0, 1

	for i := 0; i < num; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int)
	go subRoutine(10, c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
package main

import (
	"fmt"
	"time"
)

func subRoutine(quit chan bool, c chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)

	if <-quit {
		fmt.Println("Quit")
	}
}

func main() {
	c := make(chan int)
	q := make(chan bool)
	go subRoutine(q, c)
	for i := range c {
		fmt.Println(i)
	}
	q <- true
	time.Sleep(time.Second * 1)
}
*/
