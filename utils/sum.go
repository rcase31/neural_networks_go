package utils

func Sum[T Numeric](A ...T) T {
	var out T
	for i := range A {
		out += A[i]
	}
	return out
}

func VSum[T Numeric](u, v []T) []T {
	l := max(len(u), len(v))
	vsum := make([]T, l)
	for i := range Zip(u, v) {
		vsum[i] = u[i] + v[i]
	}
	return vsum
}

func SumV[T Numeric](v []T) T {
	var s T
	for _, u := range v {
		s += u
	}
	return s
}

func Mult[T Numeric](v []T, a T) []T {
	vm := make([]T, len(v))
	for i := range vm {
		vm[i] = v[i] * a
	}
	return vm
}
