package main

import (
	"fmt"
	"log"
	"sync"
)

func Go(x func()) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("%v\n", e)
			}
			wg.Done()
		}()
		x()
	}()
	wg.Wait()
}

func main() {
	Go(tryPanic)
	Go(noPanic)
	fmt.Println("来了")
}
func tryPanic() {
	panic("干他")
}
func noPanic() {
	fmt.Println("不干他了")
}
