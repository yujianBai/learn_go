package main

import(
	"fmt"
	"encoding/xml"
)

type persion struct{
	Name string `xml:"name,attr"` //将name 作为xml文档， persion标签 的一个属性, 后边的小写name 如果不给，默认就按照struct 的 Name给出名字
	Age int
}

func main(){
	per := persion{Name:"bai", Age:18}
	if data, err := xml.Marshal(per); err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println(string(data))
	}
	fmt.Println("MarshalIndent 加前缀")
	data, err := xml.MarshalIndent(per, "", " "); 
	if err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println(string(data))
	}

	fmt.Println("xml 的反序列化")
	p2 := new(persion)
	if err := xml.Unmarshal(data, p2);err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(p2)

}