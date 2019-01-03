package main

import "fmt"

func main(){
	flt_var := []float64{5.0, 6.0, 7.7, 9.0}
	sum := 0.0
	for _, v := range flt_var{
		sum += v
	}
	fmt.Println(float64(len(flt_var)))
	average := sum / float64(len(flt_var))
	fmt.Println(average)
}