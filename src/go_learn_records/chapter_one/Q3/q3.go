package main

import "fmt"

/*
1. 解决这个叫做 Fizz-Buzz[31] 的问题：
编写一个程序，打印从 1 到 100 的数字。当是三个倍数就打印 “Fizz” 代
替数字，当是五的倍数就打印 “Buzz” 。当数字同时是三和五的倍数
时，打印 “FizzBuzz” 。
*/
func FizzBuzz(num int)string{
		if (num%3)==0 && (num%5) ==0{
			return "FizzBuzz" 
		}else if(num%3) == 0{
			return "Fizz" 
		}else if(num%5) == 0{
			return "Buzz" 
		}else{
			return "0"
		}
}


func main(){
	for i := 1;i <101; i++{
		res := FizzBuzz(i)
		if res !="0"{
			fmt.Println(i, res)
		}else{
			fmt.Println(i)
		}
	}
}