package main

import "fmt"

func DuplicateEncode(word string) string {
  res := ""
  mapp := make(map[string]int)
  for i:=0; i<len(word);i++{
    mapp[string(word[i])]+=1
  }
  fmt.Println(mapp)
  for _,v := range word{
    if mapp[string(v)] == 1 {
      res += "("
    }else {
        res += ")"
    }
  }
  return res
}

func main(){
    fmt.Println(DuplicateEncode("recede"))
}