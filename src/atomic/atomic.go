package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct{
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment(){
	a.lock.Lock()
	a.value++
	defer a.lock.Unlock()
}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}


func main() {
	var a atomicInt
	a.increment()
	go func(){
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a.get())
}
