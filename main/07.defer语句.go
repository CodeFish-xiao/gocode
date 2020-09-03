package main
/*
defer 语句会将函数推迟到外层函数返回之后执行。
推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
 */
import (
	"fmt"
)

func main()  {
	defer fmt.Println("world")
	defer fmt.Println("my")
	fmt.Println("hello")
	
}