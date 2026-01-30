package main
import "fmt"

func EvenOdd(x int) string{
	if x % 2 == 0{
		return "Even"
	}
	return "Odd"
}

func main(){
	fmt.Println(EvenOdd(19))
}