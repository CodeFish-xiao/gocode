package main

import (
	"fmt"
	"math"
)
//接口I中有M方法
type I interface {
	M()
}
//类型T
type T struct {
	S string
}
//类型T实现接口I
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64
//类型F实现接口I
func (f F) M() {
	fmt.Println(f)
}
func (f F) N() {
	fmt.Println(f)
}
func main() {
	var i I//实例化接口？？？
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()//只能实例化接口方法，获取不到除了接口方法外的其他方法
	p :=F(math.Pi)
	p.N()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
