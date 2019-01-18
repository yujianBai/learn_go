package main

import (
	"fmt"
)

func testSlice(){
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice = arr[2:5]
	fmt.Println("slice", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
	slice = arr[2:4]
	fmt.Println("slice", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
	slice = arr[:]
	fmt.Println("slice", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func testSlice2_show_add(){
	arr_1 := [5]int{1, 2, 3, 4, 5}
	slice_1 := arr_1[1:]

	fmt.Printf("%p\n", slice_1)
	fmt.Printf("%p\n", &arr_1[1])
}
func main(){
	// testSlice()
	testSlice2_show_add()
}