package utils

func DotProduct[T Numeric](A, B []T) []T {
	longerLen := max(len(A), len(B))
	out := make([]T, 0, longerLen)
	for i := range Zip(A, B) {
		m := A[i] * B[i]
		out = append(out, m)
	}
	return out
}

func SumProduct[T Numeric](A, B []T) T {
	v := DotProduct(A, B)
	return SumV(v)
}
