package main

import (
	"fmt"
)

func main(){
	x := []int{1, 2, 3}

	i := 2

	//switch N case ? 是比较N和？是否相等， 相等进入case 否则进入default
	//fallthrough 关键字是，当case比配到了条件时， 进入下一case
	switch i{
	case x[0]:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
		fallthrough
	default:
		fmt.Println("C")
	}
}