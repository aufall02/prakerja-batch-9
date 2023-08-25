package main

import "fmt"

func main() {

	//anonymus function type 1
	func(){
		fmt.Println("anonymus function")
	}()

	//anonymus function type 2 store in variable
	prin :=  func(){
		fmt.Println("anonymus function")
	}
	prin()

}

// declare function
// type 1
func isFunction() {
	fmt.Println("ini function 1")
}

// type 2 : with params
func isFunction2(a int, b int) {
	fmt.Println("ini function 1")
}

// type 3 : with params & return
func isFunction3(a int, b int) int {
	return a + b;
}

// type 4 : with params & return lebih dari satu 
func isFunction4(a int, b int) (int, string) {
	return a, "ini function 4";
}

// type 4 : with params & return any
func isFunction5(a int, b int) interface{} {
	return "ini function 4";
}