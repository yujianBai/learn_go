package main

import (
	"fmt"
)

func jiechen(n int)int{ // 递归，问题可以拆分为问题的子问题， 问题的规模是不断缩小的
	if n == 1{
		return 1
	}
	if n == 0{
		return 0
	}
	return jiechen(n -1) * n
}

func fib(n int)int{
	if n <= 1{
		return 1
	}
	return fib(n-1) + fib(n -2)
}

func fib_list(n int){
	for i:= 1; i<=n; i++{
		fmt.Printf("%d ", fib(i))
	}
	fmt.Println()
}


func fib_list2(){
	fib_list  := []int{1, 2}
	for i := 2; i < 10; i++{
		fib_list = append(fib_list, (fib_list[i-1] + fib_list[i-2]))
	}
	fmt.Println(fib_list)
}

func main(){
	// fmt.Println(jiechen(2))
	fib_list(10)
	// fib_list2()
}