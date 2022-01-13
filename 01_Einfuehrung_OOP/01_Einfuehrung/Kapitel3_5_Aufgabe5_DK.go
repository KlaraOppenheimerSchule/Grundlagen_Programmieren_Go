package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Wie viele Pins sollen erstellt werden?")
	scanner.Scan()
	uNumStr := scanner.Text()
	var usrNum int
	usrNum, err := strconv.Atoi(uNumStr)
	if err != nil {
		fmt.Println(err)
	}
	var pin string
	for i := 0; i < usrNum; i++ {
		for j := 0; j < 4; j++ {
			pin += strconv.FormatInt(int64(rand.Intn(10)), 10)
		}
		fmt.Println(pin)
		pin = ""
	}
}
