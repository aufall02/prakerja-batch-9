package main

import (
	"fmt"
)

func main() {
	//declare
	var scoreData []int
	//tambah data
	scoreData = append(scoreData, []int{1,2,4,5,6,5,7,8,6,}...)
	//pootng 
	scoreData = scoreData[1:4]
	// scoreData1 = scoreData[4:6]
	scoreData = append(scoreData,scoreData[4:6]... )
	fmt.Println(scoreData)
	//range
	for i,v := range scoreData {
		fmt.Println(i , " => ", v)
	}

	isSlice := make([]int, 10) 
	fmt.Println(isSlice)

}