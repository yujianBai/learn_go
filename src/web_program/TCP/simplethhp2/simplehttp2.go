package main

import (
	"fmt"
	"os"
	"net"
	"bytes"
	"io/ioutil"
)


func main(){
	if len(os.Args) != 2{
		fmt.Println("Usage:%s host:port", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[1]
	fmt.Printf("service %s\n", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD/ HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.Readall(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}


func checkError(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
