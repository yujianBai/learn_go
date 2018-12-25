package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a, b int, op string) int {
	var ret int

	switch op {
		case "+":
			ret = a+b
		case "-":
			ret = a-b
		case "*":
			ret = a*b
		case "/":
			ret = a/b
		default:
			panic("unsuportped operator:" + op)
	}
	return ret
}

func grade(score int) string  {
	g :=""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func main(){
	const filename  = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("err : %s",err)
	}else{
		fmt.Println("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(70),
		grade(100),
		eval(1, 2, "+"))

}




