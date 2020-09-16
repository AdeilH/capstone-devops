package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//dogPic is struct
type dogPic struct {
	Message string `json:"message"`
	Status  string `json:"string"`
}

type dogDetails struct {
	URL   string
	Breed string
}

func getPicURL() string {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	randomDog := dogPic{}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonErr := json.Unmarshal(body, &randomDog)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return (randomDog.Message)
}

func serveTemplate() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template.html"))
		URL := getPicURL()
		splitURL := strings.Split(URL, "/")
		dog := dogDetails{}
		dog.Breed = splitURL[4]
		dog.URL = URL
		tmpl.Execute(w, dog)
	})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8000"
	}
	serveTemplate()
	http.ListenAndServe(port, nil)
}
