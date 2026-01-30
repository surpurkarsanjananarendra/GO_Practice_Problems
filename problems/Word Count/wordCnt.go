package main

import ("fmt")   

func wordCnt(s string) map[string]int{
	m := make(map[string]int)
	res := ""
	for i := 0; i < len(s); i++ {
		a := []rune(s)[i]
		fmt.Println(a)                 // ...prints ascii valuesof char
		if s[i] == ' ' {
			m[res]++
			res = ""
		}else {
			res = res + string(s[i])
		}
	}
	m[res]++
	return m
}

func main(){
	fmt.Println(wordCnt("I am learning GO!"))
}