package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Println("FIZZBUZZ!")
		} else if i%3 == 0 {
			fmt.Println("FIZZ!")
		} else if i%5 == 0 {
			fmt.Println("BUZZ!")
		} else {
			fmt.Println(i)
		}
	}
}
