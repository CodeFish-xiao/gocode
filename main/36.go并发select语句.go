package main
import "fmt"
//select 语句使一个 Go 程可以等待多个通信操作。
//select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
func Myfibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:	// 如果成功向c写入数据，则进行该case处理语句
			x, y = y, x+y
		case <-quit:// 如果q信道成功读到数据，则进行该case处理语句
			fmt.Println("quit")
			return
		}
	}
	/*
	//select基本用法
	select {
	case <- chan1:
	// 如果chan1成功读到数据，则进行该case处理语句
	case chan2 <- 1:
	// 如果成功向chan2写入数据，则进行该case处理语句
	default:
	// 如果上面都没有成功，则进入default处理流程
	 */
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	Myfibonacci2(c, quit)
}
