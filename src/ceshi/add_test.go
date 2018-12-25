package ceshi

import "testing"

func calcTriangle(a, b int)int{
	return a+b
}

func TestTriangle(t *testing.T){
	tests :=[]struct{a, b, c int}{
		{3, 4, 7},
		{5, 12, 17},
		{8, 15, 23},
		{12, 35, 47},
		{3000, 4000, 7000},
	}

	for _, tt := range tests{
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c{
			t.Errorf("calcTringle(%d, %d) got %d excepted %d",
				tt.a, tt.b, actual, tt.c )
		}
	}
}


func BenchmarkSubstr(b *testing.B){
	a := 3000
	d := 4000
	c := 7000
	for i:=0; i< b.N; i++{
		if actual := calcTriangle(a, d); actual != c{
			b.Errorf("calcTringle(%d, %d) got %d excepted %d",
				a, d, actual, c )
		}
	}
}
