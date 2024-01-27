package utils

type Numeric interface {
	~int | ~int64 | ~int32 | ~float32 | ~float64
}

type Array interface {
	[]any
}
