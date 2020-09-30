package main

import (
	"fmt"
	"math"
)

/*
函数也是值。它们可以像其它值一样传递。
函数值可以用作函数的参数或返回值。
*/
/**
参数为fn func(float64, float64) float64格式类型函数方可以进行，而不能只是一个返回值
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4) //这里的返回值将传进来的函数执行了一遍，得到最后的结果进行了返回
}

func main() {
	hypot := func(x, y float64) float64 { //hypot的值为这个函数的返回值：return math.Sqrt(x*x + y*y)
		return math.Sqrt(x*x + y*y)
	}
	test1 := func(n, m float64) float64 {
		return 1
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(test1))
	fmt.Println(compute(math.Pow))
}
