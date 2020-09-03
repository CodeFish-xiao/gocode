package main

import (
	"fmt"
	"math"
)

//接口类型 是由一组方法签名定义的集合。
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := MyVertex2{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = &v

	fmt.Println(a.Abs())
}

type MyFloat1 float64
// 此方法表示类型MyFloat1  实现了接口Abser，但我们无需显式声明此事。
func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyVertex2 struct {
	X, Y float64
}

func (v *MyVertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
