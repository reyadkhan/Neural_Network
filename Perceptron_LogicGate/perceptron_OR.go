package main

import (
	"fmt"
	"math/rand"
)

type Perceptron struct {
	Input []int32
	Output int32
	Weight []float32
	Bias float32
}

func activationFunc(weightedSum float64) int32  {
	if weightedSum >= 1 {
		return 1
	} else {
		return 0
	}
}

func (p Perceptron) perceptron() int32 {
	var weightedSum float64
	for i := 0; i < len(p.Input); i++ {
		weightedSum += float64(p.Weight[i]) * float64(p.Input[i])
	}
	weightedSum += float64(p.Bias)
	return activationFunc(weightedSum)
}

func (p Perceptron) errorCalc(predResult int32) float32 {
	return float32(p.Output - predResult)
}

func main()  {
	bias := float32(1)
	learningRate := float32(0.1)
	var resultCount int32
	andGateInputs := [][]int32{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	andGateOutputs := []int32{0, 1, 1, 1}
	var iterationCount int
	inputWeight_1 := rand.Float32()
	inputWeight_2 := rand.Float32()
	for resultCount < 4 {
		fmt.Println("Iteration number = ", iterationCount)
		for i := 0; i < len(andGateInputs); i++ {
			percept := Perceptron{[]int32{andGateInputs[i][0], andGateInputs[i][1]}, andGateOutputs[i], []float32{inputWeight_1, inputWeight_2}, float32(bias)}
			pred_result := percept.perceptron()
			fmt.Printf("Target output %v ------ Predicted output %v\n", percept.Output, pred_result)
			error := percept.errorCalc(pred_result)
			if error == 0 {
				resultCount ++
			} else {
				resultCount = 0
			}
			bias += error * learningRate
			inputWeight_1 += error * learningRate * float32(andGateInputs[i][0])
			inputWeight_2 += error * learningRate * float32(andGateInputs[i][1])
		}
		fmt.Println("*************** iration end ******************* \n")
		iterationCount ++
	}
}
