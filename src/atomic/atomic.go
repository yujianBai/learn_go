package main

import (
	"fmt"
	"sync"
	"time"
)
// 演示 sync.Mutex 锁 的使用
// go run -race atomic.go //查看访问冲突的数据

// go语言中的传统同步机制：
//	WaitGroup
//	Mutex
//	Cond
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
	// time.Sleep(time.Second)
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
