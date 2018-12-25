package main

import "fmt"
import "time"

func Add(x, y int){
	z := x+y
	fmt.Printf("%d + %d = %d\n",x, y, z)
}

func main(){
	for i:=0; i< 10; i++{
		go Add(i, i)
	}

	time.Sleep(time.Second) // 防止main 程序退出了， 并发程序还没执行
}

