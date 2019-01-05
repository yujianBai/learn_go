package main

import (
	"fmt"
	// "os"
	"flag"
)

func style(){
	//格式化定义
	methodPtr := flag.String("method", "default", "method of sample")
	varluePtr := flag.Int("value", -1, "value of sample")

	//解析
	flag.Parse()
	fmt.Println(*methodPtr, *varluePtr)
}

func style2(){
	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")
	
	flag.Parse()
	fmt.Println(method, value)
}

func main(){
	// args := os.Args
	// fmt.Printf("os.Args get : %v\n", args)

	// for index, v := range args{
		// println(index, v)
	// }



	//style 和 style2 都是可以加关键字的， 使用区别于os.Args
	fmt.Println("----------style----------")
	style()
	fmt.Println("----------style 2----------")
	// style2()
}
