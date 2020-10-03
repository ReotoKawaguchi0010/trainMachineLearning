package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func initWeight() [][]float64{
	rand.Seed(time.Now().UnixNano())
	var initWeight1 []float64   = []float64{rand.Float64(), rand.Float64(), rand.Float64()}
	var initWeight2 []float64   = []float64{rand.Float64(), rand.Float64(), rand.Float64()}
	var initWeight3 []float64   = []float64{rand.Float64(), rand.Float64(), rand.Float64()}
	var initWeightL [][]float64 = [][]float64{initWeight1, initWeight2, initWeight3}
	return initWeightL
}


func NeuralNetwork() {
	initWeightL := initWeight()
	n1 := []float64{0, 0, 0}
	n2 := []float64{1, 0, 1}
	n3 := []float64{0, 1, 1}
	n4 := []float64{1, 1, 0}

	epochs := 10000
	alpha  := 0.1

	input := [][]float64{n1, n2, n3, n4}
	w111  := initWeightL[0][0]
	w112  := initWeightL[0][1]
	b11   := initWeightL[0][2]
	w121  := initWeightL[1][0]
	w122  := initWeightL[1][1]
	b12   := initWeightL[1][2]
	w211  := initWeightL[2][0]
	w212  := initWeightL[2][1]
	b21   := initWeightL[2][2]

	for t := 0; t<epochs; t++{
		for i := range input{
			Y := input[i][2]
			z11 := b11 + w111*input[i][0] + w112*input[i][1]
			z12 := b12 + w121*input[i][0] + w122*input[i][1]

			o11 := sigmoid(z11)
			o12 := sigmoid(z12)

			z21 := b21 + w211*o11 + w212*o12

			o21 := sigmoid(z21)
			err := math.Pow(Y-o21, 2.0)

			dw111 := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*w211*sigmoid(z11)*(1-sigmoid(z11))*input[i][0]
			dw112 := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*w211*sigmoid(z11)*(1-sigmoid(z11))*input[i][1]
			db11  := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*w211*sigmoid(z11)*(1-sigmoid(z11))

			dw121 := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*w212*sigmoid(z12)*(1-sigmoid(z12))*input[i][0]
			dw122 := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*w212*sigmoid(z12)*(1-sigmoid(z12))*input[i][1]
			db12  := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*w212*sigmoid(z12)*(1-sigmoid(z12))

			dw211 := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*o11
			dw212 := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))*o12
			db21  := 2*(o21-Y)*sigmoid(z21)*(1-sigmoid(z21))

			w111 := w111 - alpha*(dw111)
			w112 := w112 - alpha*(dw112)
			b11  := b11  - alpha*(db11)
			w121 := w121 - alpha*(dw121)
			w122 := w122 - alpha*(dw122)
			b12  := b12  - alpha*(db12)
			w211 := w211 - alpha*(dw211)
			w212 := w212 - alpha*(dw212)
			b21  := b21  - alpha*(db21)

			fmt.Printf("w111: %s\n", w111)
			fmt.Printf("w112: %s\n", w112)
			fmt.Printf("b11: %s\n", b11)
			fmt.Printf("w121: %s\n", w121)
			fmt.Printf("w122: %s\n", w122)
			fmt.Printf("b12: %s\n", b12)
			fmt.Printf("w211: %s\n", w211)
			fmt.Printf("w212: %s\n", w212)
			fmt.Printf("b21: %s\n", b21)
			fmt.Printf("err: %s\n", err)
		}
	}
	x1 := []float64{0,0,1,1}
	x2 := []float64{0,1,0,1}

	for j := range x1{
		z11 := b11 + w111*x1[j] + w112*x2[j]
		z12 := b12 + w121*x1[j] + w122*x2[j]

		o11 := sigmoid(z11)
		o12 := sigmoid(z12)

		z21 := b21 + w211*o11 + w212*o12

		o21 := sigmoid(z21)
		fmt.Println(o21)
	}

}