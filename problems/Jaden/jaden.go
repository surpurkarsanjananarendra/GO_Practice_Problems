package main

import("strings"
		"unicode"
		"fmt")

func ToJadenCase(str string) string {
  runes := strings.Fields(str)
  res := ""
  for i := range(runes){
	res += string(unicode.ToUpper(rune(runes[i][0]))) + string(runes[i][1:]) + " "
  }
  var ans strings.Builder
	for i, r := range res {
		if i > 0 && unicode.IsUpper(r) {
			ans.WriteRune(' ')
		}
		ans.WriteRune(r)
	}
  return ans.String()
}

func main(){
	s := "I am an intern in coditas"
	fmt.Println(ToJadenCase(s))
}
