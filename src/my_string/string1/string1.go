package main

import (
	"fmt"
	"io/ioutil"
	// "os"
)

// func Name(file []os.FileInfo)string{
	// return file.name
// }

type man struct {
	name string
	addr string
}

func main(){
	int1 := 1
	fmt.Printf("%v\n", int1)

	fileInfoArr, err := ioutil.ReadDir("./")
	if err != nil{
		fmt.Println(err)
	}
	// fmt.Println(fileInfoArr)
	for _, files := range fileInfoArr{
		fmt.Println("###############")
		fmt.Println(files.Name())
		fmt.Println(files.Size())
		fmt.Println("###############.....")
		// for k, v := range files{
			// fmt.Println(v)
		// }
	} 

	var2 := []string{"b", "bbbbb", "123456"}
	// var3 := []int{1, 2, 3, 5}

	v1 := []string{}
	fmt.Println(v1)
	for _, v := range var2{
		v1 = append(v1, v)
	}
	fmt.Println(v1)


	fmt.Println("----------map")
	// var manlist map [string] man
	manlist  := make(map [string] man)
	manlist["01"] = man{"baiyu", "xi'an"}
	manlist["02"] = man{"baiyu2", "shenzhen"}
	fmt.Println(manlist["01"])
	fmt.Println(manlist["02"])
}