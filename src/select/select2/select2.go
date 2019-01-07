package main

import (
	"fmt"
)

func main(){
	var c1, c2 chan int

	for{
		select{
		case n := <-c1:
			fmt.Printf("receive from c1\n", n)
		case n := <-c2:
			fmt.Printf("receive from c2\n", n)
		default:
			fmt.Println("no value")
		}
	}
}