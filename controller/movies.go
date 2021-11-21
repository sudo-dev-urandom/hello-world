package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Search []Search `json:"Search"`
}

// A Pokemon Struct to map every pokemon to.
type Search struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type Detail struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Dvd        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

func SearchApi(s string, api string, page int) int {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", api, s, page)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Search)
	fmt.Println(len(responseObject.Search))

	for i := 0; i < len(responseObject.Search); i++ {
		fmt.Println(responseObject.Search[i].Title)
	}
	return len(responseObject.Search)
}

func GetDetail(id string, api string, page int) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s&page=%d", api, id, page)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Detail
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject)
	fmt.Println(responseObject.Year)
	fmt.Println(responseObject.Rated)
	fmt.Println(responseObject.Released)
	fmt.Println(responseObject.Runtime)
	fmt.Println(responseObject.Genre)
	fmt.Println(responseObject.Director)
	fmt.Println(responseObject.Writer)
	fmt.Println(responseObject.Actors)
	fmt.Println(responseObject.Plot)
	fmt.Println(responseObject.Language)
	fmt.Println(responseObject.Country)
	fmt.Println(responseObject.Awards)
	fmt.Println(responseObject.Poster)

	fmt.Println(len(responseObject.Ratings))

	for i := 0; i < len(responseObject.Ratings); i++ {
		fmt.Println(responseObject.Ratings[i].Source)
		fmt.Println(responseObject.Ratings[i].Value)
	}
}
