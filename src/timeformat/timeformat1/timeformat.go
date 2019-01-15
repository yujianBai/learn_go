package main

import (
	"time"
	"fmt"
)

func main(){
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006-01-02 15:04:00"))
	fmt.Println(now.Format("2006/1/02")) 
	// 格式化很恶心， 需要记住这个时间。 这个是go诞生的时间
}