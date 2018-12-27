package main

import (
	"io"
	"net/http"
	"log"
)


func helloHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello world")
}

func helloHandler2(w http.ResponseWriter, r *http.Request){
	log.Println("in nihao")
	io.WriteString(w, "你好")
}

func main(){
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/nihao", helloHandler2)	
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("ListenAndServe:", err.Error())
	}



}