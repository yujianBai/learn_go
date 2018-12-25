package main

import (
	"fmt"
	"regexp"
)

const text = `my emails is ccmouse@gmail.com
my email1 is 1ccmouse@gmail.com
my email2 is 2ccmouse@gmail.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	//match := re.FindString(text)
	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range(match){
		fmt.Println(m)
	}

}
