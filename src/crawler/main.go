package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	// "time"
)

func main() {
	engine.Run(engine.Request{
		Url : "http://www.zhenai.com/zhenghun",
		ParserFunc : parser.ParserCityList,
	})
	// time.Sleep(10 * time.Second)
}


