# Let's Go

## Info
This repo is made to track my progress and thoughts while building projects in Go. 

Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
## Table of Contents
 - [Project 1: Go-server: Build a simple Web Server with Go](#project-1-build-a-simple-web-server-with-go)

- [Project 2: CRUD Movie API](#project-2-build-a-crud-api-with-go)

## Project 1. Build a simple web server with Go
### Basic Structure:
>               --> / --> index.html
>       Server  --> /hello  --> hello func
>               --> /form --> form func --> form.html

### What is a web server?

### hello function

### form function


## Project 2. Build a CRUD API with Go

Structure of the API:
![structure of the API:](imgs/APIstructure.png)

### Install gorilla/mux:

```
go get -u "github.com/gorilla/mux"
```
 It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
### The packages that we will use in this projects include:
    - "encoding/json" //data in json and sent to postman
	- "fmt"
	- "log" 
	- "math/rand"
	- "net/http" //create a server in golang
	- "strconv"  //string converter
	- "github.com/gorilla/mux"
### Design 2 Endpoints: /movies and /movies/id

### struct is an objecct in javascript with key-value pairs 
Here, the movie struct and director struct will be associated such that every movie will have one director. 

There would be 2 properties within the director struct: 2 strings each represent the director's first name and last name:

    - Director struct properties:
        - Firstname string `json:"firstname"`
	    - Lastname  string `json: "lastname"`

There would be 4 properties within the movie struct: 3 Strings and 1 Director struct that we just made above:

    - Movie Struct properties:
	    - ID    string `json:"id"`
	    - Isbn  string `json:"isbn"`
	    - Title string `json:"title"`
	    - Director *Director `json:"director"`
### Movies array (an array of Movie struct!)
```
var movies []Movie
```
### Functions that we use to CRUD (Create, read, update, and delete):

### getMovies function
This function will return the movie with the id specified.
```go
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
```

### getMovie function

```GO
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
```

### createMovie function
```GO
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
```

### updateMovie function
```GO
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
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
```

### deleteMovies function
```go
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
```
### main function
Structure: 
1) mux.NewRouter
2) append two new movie objects into the movies array we built.
```go
movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
```
> **Note:**  append(arrayname, item)
> each object of the struct: 
> 

3) 5 differenet routes specified:
    - getMovies - "GET"
    - getMovie - "GET"
    - createMovie - "POST"
    - updateMovie - "PUT"
    - deleteMovie - "DELETE"
4) Print on screen that the server is starting at port 8000.
```GO
	fmt.Printf("Starting server at port 8000\n")
```
5) 



### Testing THE API in Postman.
