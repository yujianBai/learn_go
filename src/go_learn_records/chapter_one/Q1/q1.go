package main

import "fmt"

func main(){
	fmt.Println("q1, 1...")
	for i:=0; i<10; i++{
		fmt.Println(i +1)
	}
	fmt.Println("q1, 2...")
	


	fmt.Println("q1, 3...")
	arr_a := []int{1, 2, 3, 4, 5, 6}
	for index, value := range arr_a{
		fmt.Println(index, value, arr_a[index])
	}

}