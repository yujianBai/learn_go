package main

import "fmt"
/*
斐波那契数列以： 1,1,2,3,5,8,13,... 开始。或者用数学形式表达： 
x 1 =1;
x 2 = 1;
x n = x n−1 + x n−2 ∀n > 2 。
编写一个函数，接受 int 值，并给出这个值得到的斐波那契数列。
*/

func fiblist(var_fib int)([]int){
	var fib_list = []int{}
	if var_fib ==1{
		fib_list = []int{1}
	}else if var_fib == 2{
		fib_list = []int{1, 1}
	}else{
		fib_list = []int{1, 1}
		for i :=2; i != var_fib; i++{
			fib_list = append(fib_list, (fib_list[i-1] + fib_list[i-2]))
		}
	}
	return fib_list
}

func fibonacci(value int)[]int{
	x := make([]int, value)
	x[0] = 1
	if value > 1{
		x[1] = 1
		for n := 2; n<value; n++{
			x[n] = x[n-1]+ x[n-2]
		}
	}
	return x
}

func main(){
	fmt.Println(fiblist(2))
	fmt.Println(fibonacci(2))
	// fmt.Println(make([]int, 2)) // make([]int, 2) 创建一个长度为2的 int 型数组
}