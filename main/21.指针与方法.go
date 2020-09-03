package main
import (
	"fmt"
	"math"
)

type V3 struct {
	X, Y float64
}

func (v V3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *V3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
//指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
//试着移除第 15 行 Scale 函数声明中的 *，观察此程序的行为如何变化。
//若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。（对于函数的其它参数也是如此。）Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值。
func main() {
	v := V3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
