package main


/*
	chan 之所以做成这样， 是有理论基础的，
	Communication Sequential Process (CSP) 模型

	chnn 应用：
		Don't communicate by sharing memory; share momory by communicating.

		不要通过共享内存来通信， 通过通信来共享内存

		(通过共享内存来通信： 我们这只一个共享的标记flag, 通过flag 的值来判断事情是否做完了)
*/
import (
	"fmt"
	"time"
)

func worker(id int,  c chan int){
	for n:= range c{ //c 有可能是空, 这里的接收可以用for range 的方式， 也可以用，varuel, ok 接收两个值
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func chanDemo(){
	//var c chan int     //这样定义时 c == nil
	var channels [10]chan int

	for i:=0; i<10; i++{
		channels[i] = make(chan int)
		go worker(i, channels[i])

	}

	for i:=0; i<10; i++{
		channels[i] <- 'a' + i
	}

	for i:=0; i<10; i++{
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel(){
	c := make(chan int, 3) // 缓冲大小为3 的chan
	go worker(0, c)
	c <- 'a' +1
	c <- 'a' +2
	c <- 'a' +3
	c <- 'a' +4
	time.Sleep(time.Millisecond)
}

func channelClose(){
	c := make(chan int)
	go worker(0, c)
	c <- 'a' +1
	c <- 'a' +2
	c <- 'a' +3
	c <- 'a' +4
	close(c) // 关闭chan后， 接收方收到的是空串
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
