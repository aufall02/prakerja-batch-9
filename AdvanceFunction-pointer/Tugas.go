package main

import (
	"fmt"
)

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	carSedan := Car{
		Type:   "sedan",
		fuelln: 1.5,
	}

	studen := Student{
		names:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		scores: []float32{80, 75, 70, 75, 60},
	}

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println(carSedan.mileage())
	fmt.Println(studen.average())
	fmt.Println(studen.max())
	fmt.Println(studen.min())
	fmt.Println(min, max)

}

// soal pertama
type Car struct {
	Type   string
	fuelln float32
}

func (car Car) mileage() (string, float32, string) {
	// (1.5 L / mill)
	jarakTempuh := car.fuelln / 1.5
	return "jarak yang bisa ditempuh oleh mobile " + car.Type + " adalah ", jarakTempuh, " mill"
}

// soal kedua
type Student struct {
	names  []string
	scores []float32
}

type minMax struct {
	name  string
	score int
}

type Skor interface {
	average() float32
	min() minMax
	max() minMax
}

func (s Student) average() float32 {
	var average, tempScore float32
	for _, score := range s.scores {
		tempScore += score
	}
	average = tempScore / float32(len(s.scores))
	return float32(average)
}

func (s Student) min() minMax {
	var min = s.scores[0]
	var sMin int

	for i, score := range s.scores {
		if score <= min {
			min = score
			sMin = i
		}
	}

	m := minMax{
		name:  s.names[sMin],
		score: int(min),
	}

	return m
}

func (s Student) max() minMax {
	var max = s.scores[0]
	var sMax int

	for i, score := range s.scores {
		if score >= max {
			max = score
			sMax = i
		}
	}

	m := minMax{
		name:  s.names[sMax],
		score: int(max),
	}

	return m
}

// Soal ketiga
func getMinMax(numbers ...*int)(min int, max int) {
	//code

	base := *numbers[0]
	
	for _, number := range numbers {
		temp := *number

		if temp <= base  {
			min = temp
		}else if temp >= base{
			max = temp
		}
	}
	
	return min, max
}
