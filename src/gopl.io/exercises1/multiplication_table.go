package main

import (
	"fmt"
)


func multiplicationTable(){
	for i:=1; i<10; i++{
		for j:=1; j<=i; j++{
			fmt.Printf("%d * %d = %d\t",i, j, i*j)
		}
		fmt.Println()
	}
}

func judge_huiwen(str string){
	bytes := []rune(str) // 处理中文时，注意使用这种方式
	for from , to := 0 , len(bytes) -1 ; from < to ; from , to = from + 1, to -1{
		bytes[from] , bytes[to] = bytes[to] , bytes[from]
	}
	tmp_str := string(bytes)
	if tmp_str == str{
		fmt.Printf("%s is huiwen\n", str)
	}else{
		fmt.Printf("%s is not huiwen\n", str)
	}
}

func judjeHuiNum(strnum string){
	str_num := []rune(strnum)

	for i :=0; i<len(str_num); i++{
		if str_num[i] != str_num[len(str_num) -i -1]{
			fmt.Printf("%s is not hui num\n", string(str_num))
			return
		}
	}

	fmt.Printf("%s is hui num\n", string(str_num))
}

func morePara(args ... int){
	fmt.Println(args)
}


func main(){
	fmt.Println("multiplicationTable")
	multiplicationTable()
	fmt.Println("---------------------hui xing shu judge")
	judjeHuiNum("323")
	judjeHuiNum("324")
	judjeHuiNum("3245423")
	judjeHuiNum("3244")
	judjeHuiNum("你好好你")

	fmt.Println("----------------------mort paras")
	morePara(1, 2, 3, 4, 5)

	fmt.Println("----------------------huiwen 2")
	judge_huiwen("123qwe")
	judge_huiwen("123321")
	judge_huiwen("你好好你")
	fmt.Println("--------------------test------------------")
	str_tmp := "hello 你好"
	fmt.Println([]byte(str_tmp))
	fmt.Println([]rune(str_tmp))

	str_tmp = "hello"
	fmt.Println([]byte(str_tmp))
	fmt.Println([]rune(str_tmp))

}
