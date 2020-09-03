package main
/*
切片文法类似于没有长度的数组文法。
这是一个数组文法：
[3]bool{true, true, false}
下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：
[]bool{true, true, false}
 */
import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}//创建数组p
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}//创建数组r
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{//这里的两个括号不可回车，自动将下面括号中的值与上述相对应
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//数组切片的默认上下界切片，下界的默认值为 0，上界则是该切片的长度。
	s1 := s[1:4]
	fmt.Println(s1)

	s2 := s[:2]
	fmt.Println(s2)

	s3 := s[1:]
	fmt.Println(s3)

	/*
	切片的长度与容量变化
	 */

	ss := []int{2, 3, 5, 7, 11, 13}
	printSlice(ss)

	// 截取切片使其长度为 0,切片是镜像投影不影响被投影的原数组，所以可以进行对于切片的扩容，但是不能超过底层数组的值
	s = s[:0]
	printSlice(ss)

	// 拓展其长度
	s = s[:4]
	printSlice(ss)

	// 舍弃前两个值
	s = s[2:]
	printSlice(ss)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}