package main

import (
	"fmt"
	"strings"
)

func betterSolution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func solution(str, ending string) bool {
	str = strings.Trim(str, " ")
	ending = strings.Trim(ending, " ")
	if len(str) == 0 && len(ending) == 0 {
		return true
	} else {
		if string(str[len(str)-1]) != string(ending[len(ending)-1]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solution("ac", "abc"))
	fmt.Println(betterSolution("ac", "abc"))
}
