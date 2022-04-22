package main

import (
	"fmt"
	"sync"
	"time"
)

type Repo interface {
	Get() (string, error)
	Set() error
}

type repoImpl struct {
	sync.Mutex
}

func (r *repoImpl) Get() (string, error) {
	r.Lock()
	defer r.Unlock()
	time.Sleep(10 * time.Second)
	return "", nil
}

func (r *repoImpl) Set() error {
	r.Lock()
	defer r.Unlock()
	time.Sleep(3 * time.Second)
	return nil
}

func NewRepo() Repo {
	return &repoImpl{}
}
func main() {
	rr := NewRepo()
	go func() {
		err := rr.Set()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err := rr.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

}
