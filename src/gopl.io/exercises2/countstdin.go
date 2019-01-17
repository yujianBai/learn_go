package main

import (
	"fmt"
	"bufio"
	"os"
)

func countNUm(str string)(engworld, numcount, spacecount , othercount int){
	t:=[]rune(str)
	for _, v := range(t){
		switch {
		case v > 'A' && v < 'z':
			engworld += 1
		case v == ' ':
			spacecount ++
		case v > '0' && v < '9':
			numcount ++
		default:
			othercount ++
		}
	}
	return
}

func main(){
	reader :=	bufio.NewReader(os.Stdin)
	ret, _, err := reader.ReadLine()
	if err != nil{
		fmt.Println("read from stdio err:", err)
		return
	}
	e, n, s, o := countNUm(string(ret))
	fmt.Println(e, n, s, o)
}