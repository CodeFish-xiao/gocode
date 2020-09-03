package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	s1 :=strings.Fields(s)//根据空格切割字符串返回字符串数组
	m:=make(map[string]int)
	for i:=0;i< len(s1);i++ {
		//v, ok := m[s1[i]]
		//fmt.Println("The value:", v, "Present?", ok)
		//fmt.Println(s1[i])
		//fmt.Println(m[s1[i]])
		m[s1[i]]=m[s1[i]]+1
	}
	return m
}
func main() {
		fmt.Println(WordCount("I am learning Go!"))
}