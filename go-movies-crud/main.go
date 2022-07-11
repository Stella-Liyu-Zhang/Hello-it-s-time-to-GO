package main

import(
	"fmt"
	"log"
	"encoding/json" //data in json and sent to postman
	"math/rand"
	"net/http" //create a server in golang
	"strconv" //string converter
	"github.com/gorilla/mux"
)

//struct is an object in javascript with key-valueu pairs 
//movie and director will be associted such that every movie will have one director.
type Movie struct{
	ID string `json:"id"`
}

type Director{

}
