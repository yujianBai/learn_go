package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() IntGen {
	a, b := 0, 1
	return func() int{ // 闭包的强大！！！
		a, b = b, a+b
		return a
	}
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

// golang 中只要是一个类型， 就可以实现接口!!!



func pritnFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	pritnFileContents(f)
	/*
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())*/
}
