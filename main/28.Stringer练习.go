package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

//golang里边 string的概念其实不是以前遇到\0结尾的概念了，
//他其实就是一块连续的内存，首地址+长度，上面那样赋值，如果p里边有\0，他不会做处理这个时候，
//如果再对这个string做其他处理就可能出问题了，比如strconv.Atoi转成int就有错误，解决办法就是需要自己写一个正规的转换函数：
func byteString(p []byte) string {//[]byte转String函数
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string{
	var s = ""
	for i:=0;i< len(ip);i++ {
		if  i<len(ip)-1{
			s=s+strconv.Itoa(int(ip[i]))+"."
			//fmt.Println(s)
		} else {
			s =s+strconv.Itoa(int(ip[i]))
		}
	}
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String())
	}
}
