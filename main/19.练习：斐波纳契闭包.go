package main
import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func(i int) int {
	a:=0//这两个值会被下面调用的函数调用
	b:=0
	return func(i int) int {
		if i==0 {
			a=0
			return a
		}else if i==1 {
			b=1
			return b
		}else if i%2==0 {
			a=a+b
			return a
		} else{
			b=a+b
			return b
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
