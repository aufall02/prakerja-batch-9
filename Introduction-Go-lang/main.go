package main

import "fmt"

func main() {

	bilPrima := 15
	bilMultiple := 21
	a := 5
	b := 10
	t := 25

	fmt.Println(sayHello("aufal"))
	fmt.Println(bilPrima, `adalah bilangan prima ? `, isPrima(bilPrima))
	fmt.Println(`adalah kelipatan `, bilMultiple, ` ? `, isMultiple7(bilMultiple))
	fmt.Println(`luas trapesium : `, areaOfTrapezoid(a, b, t))
}

func sayHello(nama string) string {
	return "Hi, " + nama + "!"
}

/*
Aturan bilangan primacd h
1. bilangan prima memiliki hanya bisa dibagi 1 dan dirinya sendiri
*/

func isPrima(bil int) bool {
	// cek apakah bilangan lebih kecil sama dengan 1
	if bil <= 1 {
		return false
	}
	// cek apakah bilangan == 2 atau bilangan == 3
	if bil == 2 || bil == 3 {
		return true
	}
	//lakukan looping untuk cek apakah bilangan bisa dibagi bilangan selain 1 dan bilangan itu sendiri
	for i := 3; i < bil; i++ {
		if bil%i == 0 {
			return false
		}
	}
	return true

}

/*
Aturan untuk ceck kelipatan bilangan
1. bilangan yang habis bila dibagi bilangan yang mau dicek
*/

func isMultiple7(bil int) bool {
	// Ceck apakah bilangan yang diinput habis dibagi 7
	return bil%7 == 0
}

//rumus luas trapesium 1/2 * (a + b) * t

func areaOfTrapezoid(a int, b int, t int) int {
	return (a + b) / 2 * t
}
