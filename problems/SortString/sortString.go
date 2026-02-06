package main 

import ("fmt"
		"strings"
		"slices"
	)

func sortString(s1, s2 string) string {
	s := removeDuplicates(string(s1) + string(s2))
    r := []rune(s)
	slices.Sort(r)
	return string(r)
}

func removeDuplicates(str string) string{
	seen :=make(map[rune]bool)
	result := strings.Builder{}

	for _, val := range str{
		if !seen[val]{
			seen[val] = true
			result.WriteRune(val)
		}
	}
	return result.String()
}

func main(){
	fmt.Println(sortString("sanjana","surpurkar"))
}