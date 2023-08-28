package main

import "fmt"

func main() {

	carSedan := Car{
		Type:   "sedan",
		fuelln: 1.5,
	}

	fmt.Println(carSedan.mileage())
}

type Car struct {
	Type   string
	fuelln float32
}

func (car Car) mileage() (string, float32, string) {
	// (1.5 L / mill)
	jarakTempuh := car.fuelln / 1.5
	return "jarak yang bisa ditempuh oleh mobile " + car.Type + " adalah ", jarakTempuh, " mill"
}
