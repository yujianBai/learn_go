package main

import (
	"fmt"
)

func aaa()(func()){

	var i int = 10
	return func(){
		fmt.Printf("i, j:%d, %d\n", i, j)
	}
}

func main(){
	var j int = 5
	a := aaa()
	a()
	// j *=2
	// a()
}