package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("You chose: ", fuel)
	var fuelChosen Fuel
	if fuel == "Diesel" {
		fuelChosen = station.fuel1
	} else if fuel == "Benzin" {
		fuelChosen = station.fuel2
	}
	fmt.Println("price/L: ", fuelChosen.getPrice())
	price := fuelChosen.getPrice() * amount
	if amount >= 100 {
		fmt.Println("You got a discount of 5%")
		price = price * 0.95
	}

	return price
}

func main() {
	var amount float32
	var a64 float64
	input := bufio.NewScanner(os.Stdin)

	station := GasStation{Fuel{1.54, "Diesel"}, Fuel{1.64, "Benzin"}}
	fmt.Print("Choose your fuel: ")
	input.Scan() // which Fuel
	f := input.Text()
	fmt.Print("Your amount: ")
	_, err := fmt.Scanf("%f", &a64)

	if err != nil {
		fmt.Println(err)
	}

	amount = float32(a64)

	fmt.Print("You have to pay: ", station.refuel(f, amount))
}
