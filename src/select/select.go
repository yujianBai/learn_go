package main

import (
	"fmt"
	"math/rand"
	"time"
)


func genrator()chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500))*time.Millisecond)
			out <-i
			i++
		}
	}()
	return out
}

func worker(id int,  c chan int){
	for n:= range c{
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int)chan <- int{
	c := make(chan int)
	go worker(id, c)
	return c
}


func main() {
	var c1, c2 chan int
	c1, c2 = genrator(), genrator()
	var worker = createWorker(0)
	var values []int

	tm := time.After(10 * time.Second) 
	tick := time.Tick(time.Second) // 定时
	for{
		var activeWorker chan <- int
		var activeValue int
		if len(values) > 0{
			activeWorker = worker
			activeValue = values[0]
		}

		select{ 
			// 非阻塞 获取chan 中的数据  
			//如果需要同时处理多个 channel，可使⽤用 select 语句。它随机选择⼀一个可⽤用 channel 做 收发操作，或执⾏行 default case
	        case n := <-c1:
	        	values = append(values, n)
		    case n := <-c2:
		    	values = append(values, n)
		    case activeWorker <- activeValue:
		    	values = values[1:]
			case <- time.After(800 * time.Millisecond):
				fmt.Println("time out")
			case <- tick:
				fmt.Println("queue len=", len(values))
			case <- tm:
				fmt.Println("bey")
				return
	    }
	}

}
