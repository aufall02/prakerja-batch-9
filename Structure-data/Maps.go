package main

import "fmt"

func main() {
	//key : value
	pendapatanBulanan := map[string]int{}
	pendapatanBulanan["senin"] = 10000
	pendapatanBulanan["selasa"] = 30000
	pendapatanBulanan["rabu"] = 30000


	result, isFound := pendapatanBulanan["selasa"]

	isMpas := make(map[string]int) 

	fmt.Println(result, isFound)
	fmt.Println(isMpas)
	for key, value := range pendapatanBulanan {
		fmt.Println(key, " => ",value)
	}
}