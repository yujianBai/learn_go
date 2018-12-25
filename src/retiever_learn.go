package main

import (
	"five_th"
	"fmt"
)
const url = "www.imooc.com"

type Retriver interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string)string
}

//接口的组合
type RetriverPoster interface {
	Retriver
	Poster
}

func session(s RetriverPoster)string{
	s.Post(url, map[string]string{
		"contents":"another fake imooc.com",
	})
	return s.Get(url)
}

func download(r Retriver) string{
	return r.Get(url)
}

func main(){
	var r Retriver
	retriever := five_th.Retrivers{Contests:"this is a fake imooc.com"}
	r = &retriever
	fmt.Printf("%T %v\n", r, r)// %T 为类型， %v 是 value
	fmt.Println(download(r))

	fmt.Println("session")
	fmt.Println(session(&retriever))
}
