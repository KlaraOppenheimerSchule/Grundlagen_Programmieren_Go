package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	check(draw(), guess())
}

func check(nums [6]int, guesses [6]int) {
	fmt.Println("Draw: ", nums)
	fmt.Println("Your guess: ", guesses)
	if nums == guesses {
		fmt.Println("Jackpot!")
	} else {
		fmt.Println("Too bad!")
	}
}

func draw() [6]int {
	var nums [6]int
	for i := 0; i < 6; i++ {
		nums[i] = rand.Intn(49) + 1
	}
	return nums
}

func guess() [6]int {
	var guesses [6]int
	for i := 0; i < 6; i++ {
		fmt.Print("Please give your guess: ")
		var input int
		var err error
		for {
			var scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("No int")
			}
			if input < 1 || input > 49 {
				fmt.Println("Please give only integer values between 1 and 49!")
			} else {
				break
			}
		}
		guesses[i] = input
		fmt.Println()
	}
	return guesses
}
