package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"os"
	"io"
)

func main(){
	TestFile :="123.txt"
	infile, inerr := os.Open(TestFile)
	fmt.Println(infile)
	if inerr == nil{
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x    %s\n", md5h.Sum([]byte("")), TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x    %s\n", sha1h.Sum([]byte("")), TestFile)
	}else{
		fmt.Println(inerr)
		os.Exit(1)
	}
}