package main

import (
	"fmt"
)

func main() {

	fmt.Println(addition(3.4, 4.1))
	fmt.Println(square(2))
	fmt.Println(is_positive(3))
	fmt.Println(is_positive(-1))

	fmt.Println(faculty(4))
}

func addition(x float64, y float64) float64 {

	return x + y

}

func square(x float64) float64 {
	return x * x
}

func is_positive(x int) bool {
	if x > 0 {
		return true
	} else {
		return false
	}
}

func faculty(x int) int {

	if x-1 == 0 {
		return x
	}

	return x * faculty(x-1)

}
