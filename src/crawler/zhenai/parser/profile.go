package parser

import (
	// "fmt"
	"log"
	"crawler/engine"
	"regexp"
	"crawler/model"
	"strconv"
)

// const citylistRe =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)".[^>]*>+([^<]+)</a>`
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]+)*.</div>`)
// var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>56岁</div>`)
var heaghRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)


func ParserProfile(contents []byte, name string) engine.ParserResult {
	log.Printf("in ParserProfile ....................")
	profile := model.Profile{}
	profile.Name = name
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil{
		profile.Age = age	
	}
	
	if height, err := strconv.Atoi(extractString(contents, heaghRe)); err != nil{
		profile.Height = height
	}

	result := engine.ParserResult{
		// Requests []Request
		Items : []interface{}{profile},
	}
	log.Printf("in parserprofile result", result, profile.Age, profile.Height)
	return result
}

func extractString(contents []byte, re *regexp.Regexp)string{
	log.Println("extracString content:", contents)
	match := re.FindSubmatch(contents)
	log.Print("in extracString:", match)
	if len(match) > 2{
		return string(match[1])
	}else{
		return ""
	}
}