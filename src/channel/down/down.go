package main

import (
	"fmt"
	"sync"
)

//func doworker(id int,  c chan int, done chan bool){
func doworker(id int,  c chan int, wg *sync.WaitGroup){
	for n:= range c{ //c 有可能是空, 这里的接收可以用for range 的方式， 也可以用，varuel, ok 接收两个值
		fmt.Printf("Worker %d received %d\n", id, n)

		//go func(){ //结束标记也是并行的
		//	done <- true
		//}()
		wg.Done()
	}
}

type worker struct {
	in chan int
	//done chan bool
	 wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup)worker{
	w := worker{
		in : make(chan  int),
		//done : make(chan bool),
		wg : wg,
	}
	go doworker(id, w.in, wg)
	return w
}

func chanDemo(){
	var workers [10] worker
	var wg sync.WaitGroup

	for i:=0; i<10; i++{
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i:=0; i<10; i++{
		workers[i].in <- 'a' + i
	}


	for i:=0; i<10; i++{
		workers[i].in <- 'A' + i
	}

	// wait for all of them
	//for _, worker := range workers{
	//	<-worker.done
	//	<-worker.done
	//}
	wg.Wait()
}

func main() {
	chanDemo()
}
