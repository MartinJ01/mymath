package mymath

import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(x int, y float64) float64 {
	return math.Yn(x, y)
}
