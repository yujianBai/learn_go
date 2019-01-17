package main

import (
	"fmt"
)

func test(a int)int{
	defer func(){
		if err:=recover(); err != nil{
			fmt.Println(err)
		}
	}()
	b := 100 /a
	return b 
}
func main()  {
	test(0)
}