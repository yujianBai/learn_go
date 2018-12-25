package main

import (
	"fmt"
	"runtime"
	"time"
)

func main1(){
	for i:= 0; i<1000; i++{
		go func(i int){  //go 的并发调用
			for{
				fmt.Printf("hello from goroutine:%d\n", i)
				//printf 有io操作，会被操作系统控制，会被中断。协程可能会主动交出控制权
			}
		}(i)

	}
	time.Sleep(time.Millisecond)//防止main函数结束， 并发程序还在执行
}

func main(){
	var a[10]int
	for i:=0; i<10; i++{
		go func(i int){
			for{
				a[i]++ // 有可能一个协程始终执行， 交不出控制权, 程序挂在这里
				runtime.Gosched()//交出控制权, 一般不常用
			}
		}(i)
	}
	time.Sleep(time.Millisecond)//防止main函数结束， 并发程序还在执行
	fmt.Println(a)
}
