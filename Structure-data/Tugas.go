package main

import (
	"fmt"
)

func main() {
	fmt.Println("======================== TUGAS 1 ===========================================")
	fmt.Println("soal pertama : ", arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println("soal kedua   : ", arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println("soal ketiga  : ", arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println("soal keempat : ", arrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println("soal kelima  : ", arrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println("soal keenam  : ", arrayMerge([]string{}, []string{}))

	fmt.Println("============================ TUGAS 2 ========================================")
	fmt.Println("soal pertama : ", Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println("soal kedua   : ", Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println("soal ketiga  : ", Mapping([]string{}))

	fmt.Println("============================ TUGAS 3 ========================================")
	fmt.Println("soal pertama 	 : ", MunculSekali("1234123"))
	fmt.Println("soal kedua 	 : ", MunculSekali("76523752"))
	fmt.Println("soal ketiga 	 : ", MunculSekali("12345"))
	fmt.Println("soal keempat 	 : ", MunculSekali("1122334455"))
	fmt.Println("soal kelima	 : ", MunculSekali("0872504"))
}

func arrayMerge(arrayA, arrayB []string) []string {

	mergeArray := []string{}
	NewArray := []string{}
	mapChecker := make(map[string]int)

	//merge array
	mergeArray = append(arrayA, arrayB...)

	// buat perulangan pada array yang sudah di merge
	for i, v := range mergeArray {
		// cek apakah nilai dari array ada di dalam map
		_, exist := mapChecker[v]
		if !exist {
			// kalau tidak ada masukkan value dari array untuk di jadikan key di map
			mapChecker[v] = i
			// lalu masukkan nilai ke dalam aray baru
			NewArray = append(NewArray, v)
		}
	}

	return NewArray
}

func Mapping(slice []string) map[string]int {

	result := make(map[string]int)

	// lakukan perulangan pada slice
	for _, v := range slice {
		value, exist := result[v]
		// cek apakah map dengan key value dari slice ada atau tidak
		if exist {
			// jika ada maka value + 1
			result[v] = value + 1
		} else {
			// jika tidak ada maka value = 1 dan mambuat data baru dengan key = value dari slice
			result[v] = 1
		}

	}

	return result

}

func MunculSekali(angka string) []string {
	nomer := []string{}
	cek := make(map[rune]int)

	for _, v := range angka {
		value, exist := cek[v]
		if !exist {
			cek[v] = 1
		} else {
			cek[v] = value + 1
		}
	}

	for k, v := range cek {
		if v == 1 {
			nomer = append(nomer, string(k))
		}
	}

	return nomer
}