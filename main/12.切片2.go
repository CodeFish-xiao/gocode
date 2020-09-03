package main

import (
	"fmt"
	"strings"
)

/*
nil 切片
切片的零值是 nil。
nil 切片的长度和容量为 0 且没有底层数组。
 */
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {//没有null，就是nil
		fmt.Println("nil!")
	}
	//s[0]=1//nil切片为长度值为nil的空切片
	//if s==nil {
	//	fmt.Println("nil!")
	//}else {fmt.Println("不为空")
	//}
    mun11:=1
	s= append(s, mun11)//可以通过append函数将nil切片赋值化，当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
	if s == nil {
		fmt.Println("nil!")
	}else {
		fmt.Println("有值了")
	}
	// append可以一次性添加多个元素
	s = append(s, 2, 3, 4)

/*
   切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
   make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
   a := make([]int, 5)  // len(a)=5
   要指定它的容量，需向 make 传入第三个参数：func make([]T, len, cap) []T
   b := make([]int, 0, 5) // len(b)=0, cap(b)=5
   b = b[:cap(b)] // len(b)=5, cap(b)=5
   b = b[1:]      // len(b)=4, cap(b)=4
 */

	a := make([]int, 5)// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	printSlice1("a", a)

	b := make([]int, 0, 5)// len(b)=0, cap(b)=5
	printSlice1("b", b)
	b=b[:2]//可以扩容到底层数组上限
	printSlice1("b", b)
	c := b[:2]
	printSlice1("c", c)
	d := c[2:5]
	printSlice1("d", d)


	/*
	   切片可包含任何类型，甚至包括其它的切片。
	*/

	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	/*
	for 循环的 range 形式可遍历切片或映射。
	当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	 */

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {//i是数组下标,v是数组副本
		fmt.Printf("2**%d = %d\n", i, v)
	}
	/*
	可以将下标或值赋予 _ 来忽略它。
	for i, _ := range pow
	for _, value := range pow
	若你只需要索引，忽略第二个变量即可。
	for i := range pow
	 */
	//只要索引
	for i:= range pow{
		fmt.Printf("索引是：%d\n",i)
	}
	//只要值
	for _,v:=range pow{
		fmt.Printf("只要值的输出为：%d\n",v)
	}
}
func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)

}
