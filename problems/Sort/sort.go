package main
import ("sort"
		"fmt")

func Sorting(b []int) []int{
		sort.Ints(b)
		return b
}

func main(){
	x := []int{10,4,2,8,6}
	fmt.Println(Sorting(x))
}