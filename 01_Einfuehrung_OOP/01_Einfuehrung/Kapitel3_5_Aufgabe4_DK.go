package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i*j < 10 {
				fmt.Print(i*j, "  ")
			} else if i*j == 100 {
				fmt.Print(i * j)
			} else {
				fmt.Print(i*j, " ")
			}

		}
		fmt.Println()
	}

}
