package main

import (
	"fmt"
	"strconv"
	_ "strings"
)

type a struct {
	u int32
	b int32
}
type b struct {
	a
}

func main() {
	var i = int64(123123123123325745)
	fmt.Println(strconv.FormatInt(i, 10))
}
