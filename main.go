package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Usaremos `tags` para referirnos al nombre original de la variable que consumimos
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := "https://jsonplaceholder.typicode.com/posts"

	res, err := client.Get(url)

	if err != nil {
		panic(err.Error())
	}

	//Creamos este array para poder leer la estructura json
	var post []Post

	err = json.NewDecoder(res.Body).Decode(&post)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(post)
}
