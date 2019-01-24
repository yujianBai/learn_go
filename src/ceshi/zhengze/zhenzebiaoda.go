package main

import (
	"regexp"
	"fmt"
)

func main(){
	match_str := "abasdgbatw123对"
	forRe := `[a-zA-z]+([0-9]+)*`

	re := regexp.MustCompile(forRe)

	match := re.FindAllStringSubmatch(match_str, -1)
	fmt.Println(match)
	for _, v := range(match){
		fmt.Println(v[1])
	}


	fmt.Println("test re.FindSubmatch")
	match1 := re.FindSubmatch([]byte(match_str))
	fmt.Println(match1)
	fmt.Println(string(match1[1]))

	fmt.Println("test_zhenai str ----------------------------------------")
	test_zhenai_str := `遇到一个大气、善良、正派、有共同语言的另一伴。相互尊重理解，共度余生。</span><!----></div> <div class="m-title" data-v-bff6f798>个人资料</div> <div class="m-content-box" data-v-bff6f798><div class="purple-btns" data-v-bff6f798><div class="m-btn purple" data-v-bff6f798>离异</div><div class="m-btn purple" data-v-bff6f798>56岁</div><div class="`
	var ageRe = regexp.MustCompile(`div class="m-btn purple" data-v-bff6f798>([0-9]+)*.</div>`)
	match2 := ageRe.FindSubmatch([]byte(test_zhenai_str))
	fmt.Println((match2))
	fmt.Println(string(match2[1]))

}