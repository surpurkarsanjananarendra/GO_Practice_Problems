package main

import ("fmt"
		"unicode")

func main(){
	var s string
	s = "Sanjana"
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	fmt.Println(string(runes))
}