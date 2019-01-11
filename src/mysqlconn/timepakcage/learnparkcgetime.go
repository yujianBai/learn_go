package main

import (
	"time"
	"fmt"
	"reflect"
)

func main(){
	DefaultTimeLoc := time.Local

	// loginTime, err := time.ParseInLocation("2006-01-02 15:04:05", lastLoginTime, DefaultTimeLoc)
	fmt.Println(DefaultTimeLoc)

	// t := time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC)
	// fmt.Printf("Go lanuched at %s\n", t.Local)

	timeNow := time.Now()
	fmt.Println(timeNow)
	// timenowstr := string(timeNow)
	// fmt.Println(time.Parse("yyyy-mm-dd hh:mm:ss", timeNow))

	type aaa struct{
		name string "namestr"
		age int
	}
	aaa_str := aaa{"bai", 18}
	fmt.Println("type:", reflect.TypeOf(aaa_str.name))

	t := reflect.TypeOf(aaa_str.name) 
	tkind := t.Kind()
	fmt.Println("tkind", tkind)
	fmt.Println("value:", reflect.ValueOf(aaa_str.name))


	// t := reflect.TypeOf(aaa_str)
	// v := reflect.ValueOf(aaa_str)
	// ttag := t.Elem().Field(0).Tag
	// name := v.Elem().Field(0).String()

	// fmt.Println(ttag)
	// fmt.Println(name)

}