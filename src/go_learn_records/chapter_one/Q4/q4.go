package main

import "fmt"
import "unicode/utf8"

func q4_1(){
	str := "A"
	for i := 0; i<100; i++{
		fmt.Println(str)
		str += "A"
	}

}


/*
建立一个程序统计字符串里的字符数量：
asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。提示： 看看 unicode/utf8 包。

为了解决这个问题，需要 unicode/utf8 包的帮助。首先，阅读一下文档 go doc
unicode/utf8 | less 。在阅读文档的时候，会注意到 func RuneCount(p []byte)
int 。然后，将 string 转换为 byte slice ：
str := "hello"
b := []byte(str) ← 转换，参阅第 63 页
*/

func q4_2(){
	str := "asSASA ddd dsjkdsjs dk白"
	fmt.Printf("String %s\n Length:%d, Runes:%d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str))) 
}

func q4_3(){
	str := "1234567890"
	append_str := "abc"
	tmp_str := []string{}

	fmt.Println(str)
	tmp_str = append(tmp_str, str[:4])
	tmp_str = append(tmp_str, append_str)
	tmp_str = append(tmp_str, str[len(append_str)+4:])
	fmt.Println((tmp_str))
	
}

func q4_4(){
	s := "asSASA ddd dsjkdsjs dk"
	// s := "foobar"
    a := []byte(s)// ← Again a conversion
    // Reverse a
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] //← Parallel assignment
    }
    fmt.Printf("%s\n", string(a)) 

}

func main(){
	// q4_1()
	// q4_2()
	q4_3()
	// q4_4()

}