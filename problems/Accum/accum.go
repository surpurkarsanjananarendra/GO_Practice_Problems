package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Accum(s string) string {
	runes := make([]string, len(s))
	for i := range runes {
		runes[i] = string(unicode.ToUpper(rune(s[i]))) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}
	return strings.Join(runes, "-")
}

func main() {
	s := "abcd"
	fmt.Println(Accum(s))
}
