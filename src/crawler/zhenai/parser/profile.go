package parser

import (
	// "fmt"
	// "log"
	"crawler/engine"
	"regexp"
	"crawler/model"
	"strconv"
)


var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var heaghRE = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)


func ParserProfile(contents []byte, name string) engine.ParserResult {
	// log.Printf("in ParserProfile ....................")
	profile := model.Profile{}
	profile.Name = name
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil{
		profile.Age = age	
	}
	
	if height, err := strconv.Atoi(extractString(contents, ageRe)); err != nil{
		profile.Height = height
	}

	result := engine.ParserResult{
		// Requests []Request
		Items : []interface{}{profile},
	}
	// log.Printf("in parserprofile result", result)
	return result
}

func extractString(contents []byte, re *regexp.Regexp)string{
	match := re.FindSubmatch(contents)
	if len(match) > 2{
		return string(match[1])
	}else{
		return ""
	}
}