package main

import (
	"fmt"
)

func main() {
	fmt.Println("soal pertama : ", arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println("soal kedua   : ", arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println("soal ketiga  : ", arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println("soal keempat : ", arrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println("soal kelima  : ", arrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println("soal keenam  : ", arrayMerge([]string{}, []string{}))

	fmt.Println("====================================================================")
	fmt.Println("soal pertama : ", Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println("soal kedua   : ", Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println("soal ketiga  : ", Mapping([]string{}))
}

func arrayMerge(arrayA, arrayB []string) []string {
	
	var mergeArray []string
	var NewArray []string
	mapChecker := make(map[string]int)


	//merge array
	mergeArray = append(arrayA, arrayB...)

	//cek apakah ada niali yang sama
	//membuat map
	

	// lakukan perulangan untuk memasukkan nilai array yang sudah di merge ke dalam map
	// karna key di map bersifat unik maka value di array diganti menjadi key di map agar otomatis menghapus value uang sama
	for i, v := range mergeArray {
		_, exist := mapChecker[v]
		if !exist {
			mapChecker[v] = i
			NewArray = append(NewArray, v)
		}
	}

	return NewArray
}

func Mapping(slice []string) map[string]int {
	result := make(map[string]int)


	for _, v := range slice {
		value, exist := result[v]
		if exist {
			result[v] = value+1
		}else{
			result[v] = 1
		}
		
	}

	return result

}
