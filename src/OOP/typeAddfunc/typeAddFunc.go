package main
import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool{
	return a < b
}

func main(){
	var a Integer = 2
	fmt.Println(a.Less(3))
}