package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var a = []int{52, 20, 45, 17, 9}
	var b = []int{60, 52, 44, 36, 28}

	c := sortArr(a, b)
	for i, v := range c {
		fmt.Println(i, ":", v)
	}
}

func sortArr(a, b []int) []int {

	//判断数组的长度
	al := len(a)
	bl := len(b)
	cl := al + bl

	fmt.Println(cl)
	//var c [cl]int // non-constant array bound cl
	c := make([]int, cl)

	fmt.Println(len(c))
	fmt.Println(cap(c))
	ai := 0
	bi := 0
	ci := 0

	for ai < al && bi < bl {

		if a[ai] > b[bi] {
			c[ci] = a[ai]
			ci++
			ai++
		} else {
			c[ci] = b[bi]
			ci++
			bi++
		}
	}

	for ai < al {
		c[ci] = a[ai]
		ci++
		ai++
	}
	for bi < bl {
		c[ci] = b[bi]
		ci++
		bi++
	}

	/*	for i, v := range c {
			fmt.Println(i, ":", v)
		}
	*/
	return c
}
