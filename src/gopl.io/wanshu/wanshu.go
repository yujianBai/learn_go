package main

import(
	"fmt"
)

func iswanshu(num int)bool{
	var sum int
	for i:=1; i<num; i++{
		if num %i ==0{
			sum += i
		}
	}
	return sum == num	
}

func getwanshu(num int){
	fmt.Printf("获取小于%d的完数\n", num)

	for i:=1; i<num; i++{
	    if iswanshu(i){
	    	fmt.Printf("%d is 完数\n", i)
		}
	}
}

func main(){
	var num int
	fmt.Scanf("%d", &num)
	getwanshu(num)
}