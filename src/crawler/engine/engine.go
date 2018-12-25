package engine

import (
	"crawler/fercher"
	"log"
)

func Run(seeds ... Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		log.Printf("Feting %s", r.Url)
		body, err := fercher.Fetch(r.Url)
		if err != nil{
			log.Panicf("Fecther:error fetching url:%s:%v", r.Url, err)
			continue
		}
		paserResult := r.ParserFunc(body)
		requests = append(requests, paserResult.Requests...)
		// 等价于，    append(requests, paserResult.Requests[0], paserResult.Requests[1])
		for _, item := range paserResult.Items{
			log.Printf("Got item :%v", item) //%v 是什么就是什么 不转义
		}
	}

}