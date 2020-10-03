package main

import (
	"fmt"
	"math/rand"
)






func simpleRegression(){
	rand.Seed(5)
	var w0 float64 = rand.Float64()
	var w1 float64 = rand.Float64()
	y := []float64{20, 23, 28, 30}
	x := []float64{30, 32, 40, 44}
	alpha := 0.0001
	var dw0 float64
	var dw1 float64

	data := 4
	epochs := 100000

	for t:=0; t<epochs; t++{
		dw0 = 0
		dw1 = 0
		for i:=0; i<data; i++{
			dw0 = dw0 + 2*w0 + 2*w1*x[i] -2*y[i]
			dw1 = dw1 + x[i]*(2*w1*x[i] + 2*w0 -2*y[i])
		}

		w0 = w0 - alpha*(dw0)
		w1 = w1 - alpha*(dw1)
		fmt.Println(dw0)
	}
	fmt.Println(w0)
	fmt.Println(w1)
}