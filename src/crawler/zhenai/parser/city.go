package parser

import (
	"crawler/engine"
	"regexp"
	// "log"
)

const cityRe =`<a href="(http://[a-zA-z]+.zhenai.com/u/[0-9]+)".[^>]*>+([^<]+)</a>`

func ParserCity(contents []byte)engine.ParserResult{
	// log.Println("in contents", contents)
	re := regexp.MustCompile(cityRe)
	mathches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range mathches{
		// log.Println("in city Parsercity matches", m)
		name := string(m[2])
		result.Items = append(result.Items, "user "+name)
		result.Requests = append(result.Requests, engine.Request{
							Url : string(m[1]),
							ParserFunc : func(c []byte) engine.ParserResult{
								return ParserProfile(c, name)
							},
						})
	}
	return result
}