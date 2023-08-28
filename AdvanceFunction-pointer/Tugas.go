package main

import "fmt"

func main() {

	carSedan := Car{
		Type:   "sedan",
		fuelln: 4,
	}

	fmt.Println(carSedan.mileage())
}

type Car struct {
	Type   string
	fuelln int
}

func (car Car) mileage() (string, float32, string) {
	// (1.5 L / mill)
	jarakTempuh := float32(car.fuelln) / 1.5
	return "jarak yang bisa ditempuh oleh mobile " + car.Type + " adalah ", jarakTempuh, " mill"
}
