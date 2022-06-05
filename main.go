package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct { // define movie struct
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct { // struct for director
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set header to json
	json.NewEncoder(w).Encode(movies)                  // encode and send response as JSON
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params from url
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // remove item from array
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // encode and send response as JSON
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params from url
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // encode and send response as JSON
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)   // decode request body into movie struct
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // generate random id
	movies = append(movies, movie)               // add movie to array
	json.NewEncoder(w).Encode(movie)             // encode and send response as JSON
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set content type to json
	params := mux.Vars(r)                              // get params from url
	for index, item := range movies {                  // loop through movies array
		if item.ID == params["id"] { // if movie id matches url id
			movies = append(movies[:index], movies[index+1:]...) // remove item from array
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // decode request body into movie struct
			movie.ID = params["id"]                    // set id to id from url
			movies = append(movies, movie)             // add movie to array
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // encode and send response as JSON
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Golang programming language", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Golang programming language 2", Director: &Director{FirstName: "John", LastName: "Doe"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
