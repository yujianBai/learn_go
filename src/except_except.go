package main

func test(){
	println("111")
	defer func() {
		if err:=recover(); err != nil{
			println(err.(string))
		}
	}()

	println("222")
	panic("panic error")
	println("333")
}

func main(){
	test()
}
