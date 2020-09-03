package main
import "fmt"
//一个结构体就跟C语言中的结构体差不多,也可以达到继承等效果，访问相对字段可用.访问,可以达到面向对象编程的效果

type Vertex struct {
	X int
	Y int
	V
}
type V struct {
	Z,T int
}

func main() {
	vs:=Vertex{1, 2,V{12,123}}
	fmt.Println(vs)
	fmt.Println(vs.V)
	//如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	p:=&vs
	p.X=1e9



	v1 := V{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 := V{Z: 1}  // T:0 被隐式地赋予
	v3 := V{}      // Z:0 T:0
	q :=&V{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	fmt.Println(
		v1,v2,v3,q)


}
