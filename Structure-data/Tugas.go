package main

import (
	"fmt"
)

func main() {
	fmt.Println("soal pertama  : " ,arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie","steve", "geese"}));
	fmt.Println("soal kedua    : " ,arrayMerge([]string{"sergei", "jin"}, []string{"jin","steve", "bryan"}));
	fmt.Println("soal ketiga   : " ,arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin","yoshimitsu", "alisa", "law"}));
	fmt.Println("soal keempat  : " ,arrayMerge([]string{}, []string{"devil jin","sergei"}));
	fmt.Println("soal kelima   : " ,arrayMerge([]string{"hwoarang"}, []string{}));
	fmt.Println("soal keenam   : " ,arrayMerge([]string{}, []string{}));
}

func arrayMerge(arrayA, arrayB []string) []string {
	mergeArray := make([]string, len(arrayA)+len(arrayB))
	//merge array
	mergeArray = append(arrayA, arrayB...)

	//cek apakah ada niali yang sama
	//membuat map 
	mapChecker := make(map[string]int)

	// lakukan perulangan untuk memasukkan nilai array yang sudah di merge ke dalam map
	// karna key di map bersifat unik maka value di array diganti menjadi key di map agar otomatis menghapus value uang sama
	for i, v := range mergeArray{
		mapChecker[v] = i
	}

	//setelah melakukan map ubah kembali key dari menjadi value di array baru
	//mebuat array baru
	newarray := make([]string, len(mergeArray))
	for key, values := range mapChecker{
		newarray[values] = key
	}

	return newarray
}