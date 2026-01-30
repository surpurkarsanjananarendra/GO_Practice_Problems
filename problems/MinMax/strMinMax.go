package main
import ("strconv"
"strings"
		"fmt")

func HighAndLow(in string) string {
  var res []string
  res = strings.Fields(in)
  fmt.Println(res)
  mi,_ := strconv.Atoi(res[0])
  ma,_ := strconv.Atoi(res[0])
  for _,val := range(res) {
    num,_ := strconv.Atoi(val)
    if num < mi{
      mi = num
    }
    if num > ma{
      ma = num
    }
  }
  return strconv.Itoa(ma) + " " + strconv.Itoa(mi)
}

func main(){
	fmt.Println(HighAndLow("5 -3 0 10 11"))
}

