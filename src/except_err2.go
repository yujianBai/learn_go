package main

import "errors"

var ErrDivZero = errors.New("division by zero")

func div(x, y int)(int, error){
	if y ==0{
		return 0, ErrDivZero
	}
	return x/y, nil

}

func main(){
	switch z, err :=div(10, 0);err{
	case nil:
		println(z)
	case ErrDivZero:
		println("error",err)
	}
}


/*
panic error两种异常的方式，
panic: 程序出现错误导致关键流程走不下去用panic
error: 其它的异常用error
*/
