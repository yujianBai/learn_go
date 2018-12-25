package main

import (
	"bufio"
	"fmt"
	"os"
	"resouse_management_error_handing/fib"
)

func tryDefer()int{ //defer 确保在函数结束时调用（以栈的顺序调用）
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return 0

}

func writeFile(filename string){
	file, err:= os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++{
		fmt.Fprintln(write, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
