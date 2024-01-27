package neuralnetwork

import "github.com/rcase31/nnetworks/utils"

func MakeNeuron(weights Weights, inputs Inputs, bias Bias) Neuron {
	wi := utils.SumProduct(weights.Values, inputs.Values)
	neuronValue := wi + Number(bias)
	neuron := Neuron{
		Value:             neuronValue,
		PreceedingWeights: &weights,
		PreceedingInputs:  &inputs,
		PreceedingBias:    bias,
	}
	return neuron
}
