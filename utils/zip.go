package utils

func Zip[T, Q any](t []T, q []Q) []struct{} {
	tl := len(t)
	ql := len(q)
	ml := max(ql, tl)
	return make([]struct{}, ml)
}
