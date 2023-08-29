package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://swapi.dev/api/"
)

func ConsumeApi() {
	response, _ := http.Get(baseURL+"people/1") 
	responseData, _ :=io.ReadAll(response.Body)
	defer  response.Body.Close()

	var LukeSkywalker StarsWarsPeople
	json.Unmarshal(responseData, &LukeSkywalker)

	fmt.Println("--------------  Print Field ----------")
	println(LukeSkywalker.Name)
	println(LukeSkywalker.Height)
	println(LukeSkywalker.Massa)
	println(LukeSkywalker.HairColor)
	println(LukeSkywalker.Films[0])

}

type StarsWarsPeople struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Massa     string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColot  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender string `json:"gender"`
	Films []string `json:"films"`

}
