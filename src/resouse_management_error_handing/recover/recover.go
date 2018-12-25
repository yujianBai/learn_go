package main

import (
	"fmt"
)

func tryRecover(){
	defer func(){
		r := recover()
		if err, ok:= r.(error);ok{
			fmt.Println("error occurren:", err)
		}else{
			//panic(r)
			panic(fmt.Sprintf(
				"I don't know what to do:%v", r))
		}
	}()
	//panic(errors.New("this is an error"))
	panic(123)

}

func main(){
	tryRecover()
}
