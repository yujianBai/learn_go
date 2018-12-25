package main

func test(){
	x, y := 10, 20

	defer func(para int){
		println("defer:",para,  y) // y闭包引用
	}(x)// x被复制？

	x += 10
	y += 100

	println("x=", x, "y=", y)
}

func main(){
	test()
}
