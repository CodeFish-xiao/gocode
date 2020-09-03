package main
//信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
//ch <- v    // 将 v 发送至信道 ch。
//v := <-ch  // 从 ch 接收值并赋予 v。
//（“箭头”就是数据流的方向。）
//和映射与切片一样，信道在使用前必须创建：
//ch := make(chan int)
//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
//以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		sum += v
		fmt.Println(i)
	}
	c <- sum // 将和送入 c
}

//发送者可通过 close 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：
//若没有值可以接收且信道已被关闭，那么在执行完
//v, ok := <-ch
//此时 ok 会被设置为 false。
//循环 for i := range c 会不断从信道接收值，直到它被关闭。
//注意： 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
//还要注意： 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
func Myfibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)//创建信道
	go sum(s[:len(s)/2], c)//计算前半段
	go sum(s[len(s)/2:], c)//计算后半段
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)



	//信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
	//
	//ch := make(chan int, 100)
	//仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
	//
	//修改示例填满缓冲区，然后看看会发生什么。
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch<- 3//信道满了会出现fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)


	ch1 := make(chan int, 10)
	go Myfibonacci(cap(ch1), ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
}
