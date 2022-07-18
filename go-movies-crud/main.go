package main

import (
	"encoding/json" //data in json and sent to postman
	"fmt"
	"log"
	"math/rand"
	"net/http" //create a server in golang
	"strconv"  //string converter

	"github.com/gorilla/mux"
)

// we won't be using database!
//struct is an object in javascript with key-value pairs

//movie and director will be associted such that every movie will have one director.
type Movie struct {
	//string
	ID    string `json:"id"`
	Isbn  String `json:"isbn"`
	Title string `json:"title"`
	// the struct "director" that we defined below
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json: "lastname"`
}

//an array
var movies []Movie

func getMovies(w https.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			//take that index, and take all the data
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	W.hEADER().sET("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.newDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	//delete the movie with the i.d that you've sent
	//add a new movie - the movie that we send in the body of the postman
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.newDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}
func main() {
	// a function inside the mux library
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: 438227, Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: 45455, Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	//5 different routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
