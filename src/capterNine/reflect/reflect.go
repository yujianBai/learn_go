package main

import(
	"fmt"
	"reflect"
)

func main(){
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//获取类型信息
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("kind is float64:", v.Kind() ==reflect.Float64)
	fmt.Println("kind of x is :", v.Kind())
	fmt.Println("value:", v.Float())

	//获取值类型
	/*
	fmt.Println("获取值类型")
	var x1 float64 = 3.4
	v1 := reflect.ValueOf(x1)
	v1.Set(4.1)  //改行报错， 由于CanSet()方法返回false , x1 是一个值类型数据，不可修改 
	//要修改， 需要通过引用调用的方式。
	fmt.Println(v1)
	*/

	//下面的示例小幅修改 了之前的例子，成功地用反射的方式修改了变量 x 的值
	fmt.Println("用反射方式修改变量x")
	var x2 float64 = 3.4
	p := reflect.ValueOf(&x2)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	vx2 := p.Elem()
	fmt.Println("setablity of v:", vx2.CanSet())

	vx2.SetFloat(7.1)
	fmt.Println("vx2 interface", vx2.Interface())
	fmt.Println("vx2:",vx2)

	fmt.Println("对结构的反射操作")
	type T struct{
		A int
		B string
	}
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	TypeOfT := s.Type()

	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, TypeOfT.Field(i).Name, f.Type(), f.Interface())
	}

}
