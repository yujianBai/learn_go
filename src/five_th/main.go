package five_th

import "fmt"

/*
	duck typing
		鸭子类型：变成体系中定义"像鸭子走路，像鸭子叫(长得像鸭子)，那么就是鸭子"
		描述事物的外部行为， 而非内部结构
		从使用角度来定义

		严格来说go 是结构化类型系统，类似duck typing

	python 中的duck typing

		def download(retriver):
			return retriver.get("www.imooc.com")
		运行时才知道传入的retriver有没有get
		如何约定？ 需要注释来说明接口

	go 语言中的duck typing

		实现的时候，想具有python， c++的duck typing的灵活性
		又具有java的类型检查

	go特点：
		接口由 使用者 定义

*/

type Retrivers struct{
	Contests string
}

func (r *Retrivers) Post(url string,form map[string]string) string{
	r.Contests = form["contents"]
	return "ok"
}

func (r *Retrivers) Get(url string) string{
	return r.Contests
}

func (r *Retrivers) String()string{
	return fmt.Sprintf("Retriever:{Contents=%s}", r.Contests)
}
