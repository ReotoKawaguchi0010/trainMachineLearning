package main

import (
	"math"
)


func sigmoid(z float64) float64{
	const E float64 = math.E
	return 1 / (1 + math.Pow(E, -z))
}

type Array struct {
	array [][]float64
}

func NewArray(array ...[]float64) *Array{
	var result [][]float64
	for i := range array{
		result = append(result, array[i])
	}
	return &Array{result}
}

func SumArray(array ...[][]float64) [][]float64{
	var Y []float64
	var R [][]float64
	for k := range array{
		if k == len(array)-1{
			break
		}
		for i := range array[0]{
			for j := range array[k][i]{
				y := array[k][i][j] + array[k+1][i][j]
				Y = append(Y, y)
				if j == 3{
					R = append(R, Y)
					Y = []float64{}
				}
			}
		}
	}
	return R
}


func main(){

	simpleRegression()



}



//
//x1 := []float64{0,0,1,1}
//x2 := []float64{0,1,0,1}
//
//x3 := []float64{0,0,2,2}
//x4 := []float64{0,2,0,2}
//
//x5 := []float64{4,2,7,2}
//x6 := []float64{3,2,9,2}
//
//test  := NewArray(x1, x2).array
//test2 := NewArray(x3, x4).array
//test3 := NewArray(x5, x6).array
//
//fmt.Println(SumArray(test, test2, test3))