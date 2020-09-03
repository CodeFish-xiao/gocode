package main
import "fmt"
/*
类型 [n]T 表示拥有 n 个 T 类型的值的数组。
表达式
var a [10]int
会将变量 a 声明为拥有 10 个整数的数组。
数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
 */
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	p:=&a[0]
	fmt.Println(*p)//使用*P传出地址的值，输出P只是地址的值

	//数组切片，类似py中的数组切片，左闭右开区间
	var s []int = primes[1:4]//输出primes[1]-->[3]
	fmt.Println(s)

	/*
	切片就像数组的引用
	切片并不存储任何数据，它只是描述了底层数组中的一段。
	更改切片的元素会修改其底层数组中对应的元素。
	与它共享底层数组的切片都会观测到这些修改。
	 */
	//这里的切片更像是指针区域，修改指针区域所带的值可以同时修改原数组的值
	s[2]=0;
	fmt.Println(primes)

}
