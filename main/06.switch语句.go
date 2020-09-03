package main
/*
Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
除非以 fallthrough 语句结束，否则分支会自动终止。
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
 */
import (
	"fmt"
	"runtime"//runtime包
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {//GOOS查看目标系统
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	fmt.Println("cpus:", runtime.NumCPU())//NumCPU查看有多少逻辑CPU返回值INT，
	fmt.Println("goroot:", runtime.GOROOT())//查看目标路径
	fmt.Println("os/platform:", runtime.GOOS)//查看目标系统

	fmt.Println("分割线————————————————————分割线")
//switch的求值顺序
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()//获取周末时间
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("分割线————————————————————分割线")

/*
   没有条件的 switch 同 switch true 一样。
 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
