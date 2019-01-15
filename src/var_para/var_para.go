package main

import (
	"fmt"
	"math/rand"
)

var a string
func main(){
	var num int
	ret_n, err := fmt.Scanf("%d", &num)
	fmt.Println("ret_n", ret_n)
	fmt.Println("err", err)
	fmt.Println("num", num)
	fmt.Println("start var para")	
	a = "G"
	fmt.Println(a)
	m()
}


func m(){
	a:= "r"
	fmt.Println(a)
	n()
	fmt.Println("get rand num")
	getRandNum()
}

func n(){
	fmt.Println(a)
}

func getRandNum(){
	for i:= 0; i<10; i++{
		rn := rand.Intn(100)
		fmt.Println(rn)
	}
	fmt.Println("get rang float64 num")

	for i:=0; i<10; i++{
		fmt.Println(rand.Float32())
	}
}