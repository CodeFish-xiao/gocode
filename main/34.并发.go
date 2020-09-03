package main
import (
	"fmt"
	"time"
)
//线程
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	go say("小爷的线程！！")
	time.Sleep(time.Second*6)
}
