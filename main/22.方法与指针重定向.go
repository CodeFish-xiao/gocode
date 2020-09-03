package main

import (
	"fmt"
	"math"
)
//比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：
//var v Vertex
//ScaleFunc(v, 5)  // 编译错误！
//ScaleFunc(&v, 5) // OK
//而以指针为接收者的方法被调用时，接收者既能为值又能为指针：
//var v Vertex
//v.Scale(5)  // OK
//p := &v
//p.Scale(10) // OK
//对于语句 v.Scale(5)，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。
type  MyVertex1 struct {
	X,Y float64
}

func (v *MyVertex1) Scale(f float64) {
	v.X=v.X*f
	v.Y=v.Y*f
}

func ScaleFunc(v *MyVertex1,f float64)  {
	v.X=v.X*f
	v.Y=v.Y*f
}
//同样的事情也发生在相反的方向。
//接受一个值作为参数的函数必须接受一个指定类型的值：
//var v Vertex
//fmt.Println(AbsFunc(v))  // OK
//fmt.Println(AbsFunc(&v)) // 编译错误！
//而以值为接收者的方法被调用时，接收者既能为值又能为指针：
//var v Vertex
//fmt.Println(v.Abs()) // OK
//p := &v
//fmt.Println(p.Abs()) // OK
//这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
func (v MyVertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
func AbsFunc(v MyVertex1) float64  {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
func main() {
	v:=MyVertex1{3,4}
	v.Scale(2)
	ScaleFunc(&v,10)


	p:=&MyVertex1{4,3}
	p.Scale(2)
	ScaleFunc(p,10)
	fmt.Println(v,p)
	v1 := MyVertex1{3, 4}
	fmt.Println(v1.Abs())
	fmt.Println(AbsFunc(v1))

	p1 := &MyVertex1{4, 3}
	fmt.Println(p1.Abs())
	fmt.Println(AbsFunc(*p1))

/*
总而言之，在使用函数和自带方法的时候，使用接收者的方法可以使用值和指针进行，而函数所标识的是什么，他就是什么
 */


}
