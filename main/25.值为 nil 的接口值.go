package main
//底层值为 nil 的接口值
//即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
//在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
//注意: 保存了 nil 具体值的接口其自身并不为 nil。
import (
	"fmt"
)

//I1接口
type I1 interface {
	M()
}
//T1结构
type T1 struct {
	S string
}
//T1为接收者实现I1接口
func (t *T1) M() {
	if t == nil {//可以在M接口中处理空指针，上流！！！！！！
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I1

	var t *T1//T1没有给实例化对象，为空指针，空指针为nil
	i = t
	Mydescribe1(i)
	i.M()

	i = &T1{"hello"}
	Mydescribe1(i)
	i.M()

	//空接口
	//指定了零个方法的接口值被称为 空接口：
	//interface{}
	//空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
	//空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
	var i1 interface{}
	Mydescribe2(i1)

	i1 = 42
	Mydescribe2(i1)

	i1 = "hello"
	Mydescribe2(i1)

}

func Mydescribe1(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)


}
func Mydescribe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

