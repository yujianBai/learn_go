package parser

import (
	"crawler/engine"
	"regexp"
)

const citylistRe =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)".[^>]*>+([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(citylistRe)
	matches := re.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParserResult{}
	limit := 2
	for _, m :=range(matches){
		 result.Items = append(result.Items, "City "+m[2])
		 result.Requests = append( result.Requests, engine.Request{
					Url:m[1],
					ParserFunc: ParserCity,
				 })
		limit --
		if limit == 0{
			break
		}
	}
	return result
}
