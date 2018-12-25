package main

import "fmt"

func Test(x [2]int){
	fmt.Printf("x:%p\n", &x)
	x[1] = 1000

	fmt.Println(x)
}

func array_read1(a []int){
	for i:= 0; i<len(a); i++{
		fmt.Println(a[i])
	}
}

func array_read2(a []int){
	for _, v :=range(a){
		fmt.Println(v)
	}
}

// 值拷贝通常会造成性能问题， 一般用数组指针或者 slice
func main(){
	a:=[2]int{}
	fmt.Printf("a:%p\n", &a)
	a[0] = 1
	Test(a)
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("array1 ....")
	array_read1(b)
	fmt.Println("array2 ....")
	array_read2(b)
}
