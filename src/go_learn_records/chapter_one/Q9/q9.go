package main

import "fmt"
import "strconv"

/*
首先定义一个新的类型来表达栈；需要一个数组（来保存键）和一个指向最后
一个元素的索引。这个小栈只能保存 10 个元素。
*/

type stack struct{
	i int
	data [10]int
}

func (s *stack)push(k int){
	if s.i+1 > 9{
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack)pop()int{
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main(){
	var s stack
	fmt.Println(s)
	s.push(25)
	fmt.Println(s)
	s.push(24)
	fmt.Printf("stack:%v\n", s)
	pop_var := s.pop()
	fmt.Println(pop_var)

	fmt.Println("stack String func")
	fmt.Println(s.String())
}