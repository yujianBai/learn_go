package fib


func Fibonacci() func() int {
	a, b := 0, 1
	return func() int{ // 闭包的强大！！！
		a, b = b, a+b
		return a
	}
}
