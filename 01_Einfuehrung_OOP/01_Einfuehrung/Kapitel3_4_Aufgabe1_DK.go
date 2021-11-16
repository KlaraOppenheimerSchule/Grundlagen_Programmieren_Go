package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Fuel struct {
	price float32
	name  string
}

type GasStation struct {
	fuel1 Fuel
	fuel2 Fuel
}

func (fuel Fuel) getName() string {
	return fuel.name
}

func (fuel Fuel) getPrice() float32 {
	return fuel.price
}

func (station GasStation) refuel(fuel string, amount float32) float32 {
	fmt.Print("You chose: ", fuel)
	var fuelChosen Fuel
	if fuel == "Diesel" {
		fuelChosen = station.fuel1
	} else if fuel == "Benzin" {
		fuelChosen = station.fuel2
	}
	fmt.Print("price/L: ", fuelChosen.getPrice())
	price := fuelChosen.getPrice() * amount
	if amount >= 100 {
		price = price * 0.95
	}

	return price
}

func main() {
	station := GasStation{Fuel{1.54, "Diesel"}, Fuel{1.64, "Benzin"}}
	reader := bufio.NewReader(os.Stdin)
	f, _ := reader.ReadString('\n') // welches Fuel
	a, _ := reader.ReadString('\n') // wie viel

	var s float32
	if s, err := strconv.ParseFloat(a, 32); err == nil {
		fmt.Println(s)
	}
	if s, err := strconv.ParseFloat(a, 64); err == nil {
		fmt.Println(s)
	}
	station.refuel(f, s)
}
