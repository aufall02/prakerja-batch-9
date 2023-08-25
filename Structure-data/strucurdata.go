package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Array declaration var <variable name> [<size_of_array>] <type_variable>
	var array [5]string
	fmt.Println(reflect.ValueOf(array).Kind())

	//assign and acces array elemt
	var country [2]string

	country[0] = "Indonesia" //insnert a value  to first element
	country[1] = "Malaysia"  //insnert a value  to second element
	fmt.Println(country[0])
	fmt.Println(country[1])

	//Inisialization with array literal
	odd_number := [3]int{1, 2, 3} // initialized whit value
	var event_number [4]int = [4]int{0,2,4}

	fmt.Println(odd_number)
	fmt.Println(event_number)

//infer length of array
	primes := [...]int{2,4,5,6,4,8,5,6,7,}
	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(len(primes))
}
