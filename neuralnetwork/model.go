package neuralnetwork

type Number float32

type NumberArray struct {
	Values []Number
	Dim    int
}

func newNumberArray(values ...Number) NumberArray {
	dim := len(values)
	na := make([]Number, 0, dim)
	for _, v := range values {
		na = append(na, v)
	}
	return NumberArray{
		Values: na,
		Dim:    dim,
	}
}

type Inputs NumberArray

func NewInputSet(values ...Number) Inputs {
	return Inputs(newNumberArray(values...))
}

type Weights NumberArray

func NewWeights(values ...Number) Weights {
	return Weights(newNumberArray(values...))
}

type LayerWeights []Weights

func NewLayerWeights(W ...Weights) LayerWeights {
	lw := make([]Weights, len(W))
	for i := range W {
		lw[i] = W[i]
	}
	return lw
}

type Bias Number

type Biases []Bias

type Neuron struct {
	Value             Number
	PreceedingWeights *Weights
	PreceedingInputs  *Inputs
	PreceedingBias    Bias
}

type NeuronLayer struct {
	Neurons []Neuron
	Dim     int
}

func (nl *NeuronLayer) AsInputs() Inputs {
	inputs := make([]Number, nl.Dim)
	for i := range nl.Neurons {
		inputs[i] = nl.Neurons[i].Value
	}
	return NewInputSet(inputs...)
}
