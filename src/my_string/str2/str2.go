package main

import (
	"fmt"
	"strings"
)

func judge_str_http(str string) string{
	if strings.HasPrefix(str, "http"){
		return str
	}else{
		return fmt.Sprintf("http://%s", str)
	}
}

func getStrSEindex(str1, str2 string)(int, int){
	var start, end int
	start = strings.Index(str1, str2)
	end = strings.LastIndex(str1, str2) 
	return start, end
}

func main(){
	str_test := "123"
	fmt.Println(judge_str_http(str_test))

	fmt.Println("judge str index")
	st, en := getStrSEindex("123qwerhello23", "123")
	fmt.Println(st, en)
}