package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	integers := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-generics Sums: %v and %v\n",
		SumIntegers(integers),
		SumFloats(floats),
	)

	fmt.Printf("Generics Sums: %v and %v\n",
		Sum[string, int64](integers),
		Sum[string, float64](floats),
	)

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		Sum(integers),
		Sum(floats),
	)
}

func SumIntegers(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

//func Sum[K comparable, V int64 | float64](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
