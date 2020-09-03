package main

import "fmt"

func main() {
	//for表达式外无需小括号 ( ) ，而大括号 { } 则是必须的
	sum := 0
	for i := 0; i < 10; i++ {//简单for循环
		sum += i
	}
	//可无初始化语句和后置语句
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	//当去掉分号的时候就是while
	sum3 := 1
	for  sum3 < 1000 {
		sum3 += sum3
	}
	//当去掉循环条件时就是无限循环
	sum4 := 0
	for  {
		sum4++
		if sum4>20 {
			break
		}
	}

	fmt.Println(sum)

	fmt.Println(sum2)
	fmt.Println(sum3)
	fmt.Println(sum4)
}
