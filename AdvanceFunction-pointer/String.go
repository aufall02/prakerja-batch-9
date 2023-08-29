package main

import (
	"fmt"
	"strings"
)

const (
	str       = "something"
	substring = "some"
)

func main() {
	fmt.Println("panjang string : ", lenString(str))
	fmt.Println("apakah string pertama mengandung string kedua  : ", contString(str, substring))
	fmt.Println("apakah string pertama sama dengan string kedua : ", compString(str, substring))
}

// len : panjang sebuah string
func lenString(str string) int {
	lenStr := len(str)
	return lenStr
}

// contains string
func contString(str1,str2 string) bool {
	return strings.Contains(str1, str2)
}

// Compare string
func compString(str1, str2 string) bool {
	return str1 == str2
}