package main

import "fmt"

func printthem(number ... int){
	for _, v := range number{
		fmt.Printf("%d\n", v)
	}
}


func para(args ... interface{}){
	fmt.Println(args)
}

func main(){
	printthem(1, 2, 3, 4)
	para("a", 1, "b", 2, "abc")
}