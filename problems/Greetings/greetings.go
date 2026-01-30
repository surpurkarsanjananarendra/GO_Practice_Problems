package main
import "fmt"

func greet(name string) string{
	msg := fmt.Sprintf("Hiee, %v Welcome",name)
	return msg
}

func main(){
	fmt.Println(greet("Bimla"))
}