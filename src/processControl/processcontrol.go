package main

import "fmt"
// import "strings"

func  main(){
	for i :=0; i<10; i++{
		for j:=0; j<=i; j++{
			fmt.Printf("A")
		}
		fmt.Println()
	}

	str := "hello world"
	for index, value := range(str){
		fmt.Printf("index[%d],var[%c], len(%d)\n",
			index, value, len([]byte(string(value))))
			// index, value, len(string(value)))
	}

	fmt.Println(len(str))
	i := 0
	LABLE:
		for ; i<10; i++{
			fmt.Println("in lcale", i)
			goto LABLE
		}


}