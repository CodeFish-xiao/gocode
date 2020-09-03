package main

import (
	"fmt"
	"math"
)
//Go 没有类。不过你可以为结构体类型定义方法。
//方法就是一类带特殊的 接收者 参数的函数。
//方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
//在此例中，Abs 方法拥有一个名为 v，类型为 V1 的接收者。
type V1 struct {
	X, Y float64
}

func (v V1) Abs() float64 {//给V1结构体写个Abs方法
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v V1) Abs2() float64 {//方法只是个带接收者参数的函数。
	return 4.0
}
type MyFloat float64

func (f MyFloat) Abs() float64 {//也可以为非结构体类型声明方法。只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	v2 := V1{3, 4}
	fmt.Println(v2.Abs())
	fmt.Println(v2.Abs2())//功能并没有改变

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())//不同的结构体或者非结构体能有相同命名的方法，也就是相当于类的概念

}
