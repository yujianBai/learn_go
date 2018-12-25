package main

import "fmt"

func adder() func(int)int{
	sum := 0 //自由变量
	return func(v int) int {
		sum += v
		return sum //函数体的局部变量, 以及局部变量的引用，共同组成了一个闭包的概念, (不单单返回的是函数体内部的变量， 以及这些参数的引用)
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int)iAdder{
	return func(v int)(int, iAdder){
		return base + v, adder2(base + v)
	}
}
/*
	"正统"函数式编程:
		不可变性：不能有状态， 只能有常量和函数
		函数只能有一个参数
*/

func main() {
	a := adder2(0)
	for i:= 0; i<10; i++{
		var s int
		s, a = a(i)
		fmt.Println(s)
	}
}
