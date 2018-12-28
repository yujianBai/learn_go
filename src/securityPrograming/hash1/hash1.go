package main

import(
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func main(){
	TestingString := "Hi, pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestingString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestingString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}