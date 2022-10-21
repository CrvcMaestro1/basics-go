package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var c, java bool
var i, j = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// basic 1
	fmt.Println("My favorite number is", rand.Intn(10))
	// basic 2
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	// basic 3
	fmt.Println(math.Pi)
	// basic 4
	fmt.Println(add(10, 2))
	// basic 6
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// basic 7
	fmt.Println(split(17))
	// basic 8
	fmt.Println(i, c, java)
	// basic 9
	fmt.Println(i, j, true, false)
	// basic 10
	k := 3
	fmt.Println(k)
	// basic 11
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
