package main

import (
	"fmt"
	"net/http"

	"github.com/shalini/mongodbapi/router"
)

func main() {
	fmt.Println("MongoDb API")
	r := router.Router()
	fmt.Println("Server is getting started")
	http.ListenAndServe(":5000", r)
	fmt.Println("Listening at port 5000")
}
