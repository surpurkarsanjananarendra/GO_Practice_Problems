package main

import (
    "fmt"
    "log"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)
    if _, err := fmt.Printf("%s\n",io.ReadAll(lr)); err != nil {
		log.Fatal(err)
	}
}
