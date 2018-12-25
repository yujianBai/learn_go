package main

import (
	"fmt"
	"net/http"
	"os"
	"resouse_management_error_handing/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error


func errWrapper(handler appHandler)func( //函数式编程， 函数作为参数作为输入， 输出
	w http.ResponseWriter, r *http.Request){
		return func(writer http.ResponseWriter, request *http.Request){
			err := handler(writer, request)
			if err != nil{
				code := http.StatusOK
				switch {
				case os.IsNotExist(err):
					http.Error(
						writer,
						http.StatusText(http.StatusNotFound),
						http.StatusNotFound)
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError
				}
				http.Error(writer, http.StatusText(code), code)
			}
		}
}
func main(){
	fmt.Println(0)
	http.HandleFunc("/list/",
		errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":18888", nil)
	if err != nil{
		fmt.Println(3)
		panic(err)
	}
}