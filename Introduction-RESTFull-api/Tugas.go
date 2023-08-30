package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	URL = "https://jsonplaceholder.typicode.com/posts"
)

type Data struct {
	UserId string `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ResponseDatas struct {
	Message string   `json:"status_data"`
	Data      Data `json:"datas"`
}

func Tugas() {
	e := echo.New()
	e.GET("/datas/:id", GetDataByIdController) //:id mengambil nilai yang ada di link/id (Param)
	e.GET("/datas", GetDatasController)
	e.POST("/datas", CreateDataController)
	e.PUT("/datas", UpdateDataController) //mengambil nilai id yang ada di link/id (QueriParam)
	e.DELETE("/datas/:id", DeleteDataController)
	e.Start(":8000")
}

// menampilkan data dari [link api](https://jsonplaceholder.typicode.com/posts)
func GetDatasController(c echo.Context) error {
	datas := getDatas()
	return c.JSON(http.StatusOK, datas)
}

// menampilkan data by id dari [link api](https://jsonplaceholder.typicode.com/posts)
func GetDataByIdController(c echo.Context) error {
	var result Data
	id, _ := strconv.Atoi(c.Param("id"))
	datas := getDatas()
	for _, data := range datas {
		if id == data.Id {
			result = data
		}
	}

	return c.JSON(http.StatusOK, result)
}

// menampilkan post data ke [link api](https://jsonplaceholder.typicode.com/posts)
func CreateDataController(c echo.Context) (err error) {
	//ambil json request dari user
	jsonBody := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		fmt.Println("error")
		return nil
	}
	//ubah data json ke byte
	out, _ := json.Marshal(&jsonBody)
	var jsonStr = []byte(out)

	// lakukan post post ke api
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//ambil body dari hasil requets ke api
	body, _ := io.ReadAll(resp.Body)

	//convert ke json lagi
	var data Data
	json.Unmarshal(body, &data)

	//send response ke
	return c.JSON(http.StatusCreated, data)

}

// update data data ke [link api](https://jsonplaceholder.typicode.com/posts)
func UpdateDataController(c echo.Context) (err error) {
	id := c.QueryParam("id")
	println(id)
	//ambil json request dari user
	jsonBody := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		fmt.Println("error")
		return nil
	}
	//ubah data json ke byte
	out, _ := json.Marshal(&jsonBody)
	var jsonStr = []byte(out)

	// lakukan post post ke api
	req, err := http.NewRequest("PUT", URL+"/"+id, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(URL+"/"+id)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//ambil body dari hasil requets ke api
	body, _ := io.ReadAll(resp.Body)

	//convert ke json lagi
	var data Data
	json.Unmarshal(body, &data)

	//send response ke
	return c.JSON(http.StatusCreated, data)

}

// deletedata data ke [link api](https://jsonplaceholder.typicode.com/posts)
func DeleteDataController(c echo.Context) (err error) {
	id := c.Param("id")
	 req, err := http.NewRequest("DELETE", URL+"/"+id, nil)
	 if err != nil {
		 panic(err)
	 }
	 client := &http.Client{}
	 resp, err := client.Do(req)
	 if err != nil {
		 panic(err)
	 }
	 defer resp.Body.Close()
	 body, err := io.ReadAll(resp.Body)
	 if err != nil {
		 panic(err)
	 }
	 var data Data
	json.Unmarshal(body, &data)
	message := fmt.Sprintf ("id %s berhasil dihapus",id)
	return c.JSON(http.StatusCreated, ResponseDatas{
		Message: message,
		Data: data,
	})

}

func getDatas() []Data {
	response, _ := http.Get(URL)
	responseData, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	var datas []Data
	json.Unmarshal(responseData, &datas)
	return datas
}
