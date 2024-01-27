package main

import (
	"fmt"

	"github.com/rcase31/nnetworks/neuralnetwork"
	"github.com/rcase31/nnetworks/utils"
)

func main() {
	P2DotProduct()
	P2DotProductImprovement()
}

func P2DotProduct() {
	inputs := [4]float32{1, 2, 3, 2.5}
	weights := [3][4]float32{
		{0.2, 0.8, -0.5, 1},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
	}
	biases := [3]float32{2, 3, 0.5}

	neuronOuts := [3]float32{0, 0, 0}
	for i := range utils.Zip(weights[:], biases[:]) {
		layerWeights := weights[i]
		bias := biases[i]
		var neuronOut float32

		for j := range utils.Zip(layerWeights[:], inputs[:]) {
			neuronOut += layerWeights[j] * inputs[j]
		}
		neuronOut += bias
		neuronOuts[i] = neuronOut
	}

	fmt.Println(inputs)
	fmt.Println(weights)
	fmt.Println(biases)
	fmt.Println(neuronOuts)

}

func P2DotProductImprovement() {
	i := neuralnetwork.NewInputSet(1, 2, 3, 2.5)
	w1 := neuralnetwork.NewWeights(.2, .8, -.5, 1)
	w2 := neuralnetwork.NewWeights(.5, -.91, .26, -.5)
	w3 := neuralnetwork.NewWeights(-.26, -.27, .17, .87)
	biases := [3]neuralnetwork.Bias{2, 3, 0.5}

	n1 := neuralnetwork.MakeNeuron(w1, i, biases[0])
	n2 := neuralnetwork.MakeNeuron(w2, i, biases[1])
	n3 := neuralnetwork.MakeNeuron(w3, i, biases[2])

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

}
