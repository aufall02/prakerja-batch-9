package main

import "fmt"

func main() {
	var age int = 100
	var ageAddres *int = &age


	fmt.Println(ageAddres)
	fmt.Println(&ageAddres)
}

func ubahAddres(addres *int)  {
	
}