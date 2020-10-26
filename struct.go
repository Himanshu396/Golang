package main

import (
	"fmt"
)

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	c := Car{
		Name:       "audi",
		Model:      "2020",
		Color:      "black",
		WeightInKg: 2000,
	}

	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	c.Color = "white"
	fmt.Println("Car: ", c)
}