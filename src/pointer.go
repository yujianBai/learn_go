package main
import "fmt"


/*
	值传递
	引用传递

*/

func swap(a , b *int){
	*a, *b = *b, *a
}

func swap2(a, b int)(int, int){
	return b, a
}

func main(){
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)

	a, b =swap2(a, b)
	fmt.Println(a, b)
}
