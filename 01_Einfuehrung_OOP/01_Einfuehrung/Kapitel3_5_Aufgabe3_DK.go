package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var digit float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Gebe eine Zahl ein: ")
	scanner.Scan()
	str := scanner.Text()
	digit, err := strconv.ParseFloat(str, 32)

	if err != nil {
		fmt.Println(err)
	}

	var times int
	for digit >= 0.001 {
		times++
		digit /= 2
		fmt.Println("next step: ", digit)
	}
	fmt.Println("Diveded by two ", times, " times")
}
