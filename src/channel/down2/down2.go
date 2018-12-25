package main

import (
	"fmt"
	"sync"
)

func doworker(id int,  c chan int, w worker){
	for n:= range c{ //c 有可能是空, 这里的接收可以用for range 的方式， 也可以用，varuel, ok 接收两个值
		fmt.Printf("Worker %d received %d\n", id, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup)worker{
	w := worker{
		in : make(chan  int),
		done : func(){
			wg.Done()
		},
	}
	go doworker(id, w.in, w)
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

	wg.Wait()
}

func main() {
	chanDemo()
}
