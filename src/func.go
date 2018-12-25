package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int)(q, r int){
	return a/b, a%b
}

func divfun(a, b int)(q int){
	return a/b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	op_name := runtime.FuncForPC(p).Name()
	fmt.Println("calling function %s whti args %d, %d", op_name, a, b)
	return op(a, b)
}

func main(){
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(divfun, 10, 3))

	fmt.Println(apply(
		func(a int, b int)int{
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 3))

}