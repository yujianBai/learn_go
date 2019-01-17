package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func bigNumAdd(a, b string)(result string){
	a = strings.Trim(a, " ")
	b = strings.Trim(b, " ") 
	if len(a) == 0 && len(b) == 0{
		return
	}
	var index1 = len(a) -1
	var index2 = len(b) -1
	var left int 

	for index1 >=0 && index2 >=0{
		c1 := a[index1] - '0'
		c2 := b[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10{
			left = 1
		}else{
			left = 0
		}
		index1 --
		index2 --
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
	}

	if index1 >= 0{
		c2 := a[index1] - '0'
		sum := + int(c2) + left
		if sum >= 10{
			left = 1
		}else{
			left = 0
		}
		index1 --
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
	}

	if index2 >= 0{
		c2 := b[index2] - '0'
		sum := + int(c2) + left
		if sum >= 10{
			left = 1
		}else{
			left = 0
		}
		index2 --
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
	}

	if left == 1{
		result = fmt.Sprintf("%c%s", '1', result)
	}
	return 
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	ret, _, err := reader.ReadLine()
	if err != nil{
		fmt.Println("get stdin err:", err)
		return
	}
	ret_slice := strings.Split(string(ret), "+")
	if len(ret_slice) < 2{
		fmt.Println("user age a + b")
	}
	fmt.Println("stdin is ", ret_slice[0], ret_slice[1])
	fmt.Println(bigNumAdd(ret_slice[0], ret_slice[1]))

}