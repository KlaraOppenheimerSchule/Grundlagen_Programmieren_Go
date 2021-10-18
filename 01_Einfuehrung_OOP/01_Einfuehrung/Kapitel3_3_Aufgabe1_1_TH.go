package main

import "fmt"

func main() {
	fmt.Println(add(42, 1337))
	fmt.Println(is_positive(0))
	fmt.Println(is_positive(-2))
	fmt.Println(is_positive(5))
	fmt.Println(fakultaet(3))
	fmt.Println(fakultaet_recursiv(3))
}
func add(a int, b int) int {
	return a + b
}
func square(x int) int {
	return x * x
}
func is_positive(x int) bool {
	var positive bool
	if x >= 0 {
		positive = true
	} else {
		positive = false
	}
	return positive
}
func fakultaet(x int) int {
	ret := 1
	for i := 1; i <= x; i++ {
		ret *= i
	}
	return ret
}
func fakultaet_recursiv(x int) int {
	if x < 2 {
		return 1
	} else {
		return x * fakultaet_recursiv(x-1)
	}
}
