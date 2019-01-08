package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[1:], " "))

	// varlsit := []int{1, 2, 3}
	// for _, v := range varlsit{
		// fmt.Println(v)
	// }

	fmt.Println("func name is :", os.Args[0])

	fmt.Println("print index and value of os.Args")
	for index, value := range os.Args[1:]{
		fmt.Printf("%d, %s\n", index, value)
	}

	fmt.Println("program end")
}