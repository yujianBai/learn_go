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

	for _, m :=range(matches){
		 result.Items = append(result.Items, m[2])
		 result.Requests = append(
		 	result.Requests,
			engine.Request{
					Url:m[1],
					ParserFunc: engine.NilParser,
		 		})

	}
	return result
}
