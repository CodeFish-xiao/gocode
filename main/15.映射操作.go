package main

import (
	"fmt"
)

/**
在映射 m 中插入或修改元素：
m[key] = elem
获取元素：
elem = m[key]
删除元素：
delete(m, key)
通过双赋值检测某个键是否存在：
elem, ok = m[key]
若 key 在 m 中，ok 为 true ；否则，ok 为 false。
若 key 不在映射中，那么 elem 是该映射元素类型的零值。
同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
elem, ok := m[key]
 */
func main() {
	m:=make(map[string]string)
	m["Answer"]=""
	fmt.Println("The value：",m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	//直接修改
	m["Answer"]="45"
	fmt.Println("修改后的The value：",m["Answer"])
	//删除数据
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	//双赋值检测,将键值赋予给前面的变量，ok接受是否存在的键值对信息，不存在时键值为0，ok为false
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}