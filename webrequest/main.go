package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("response of type : %T\n", response)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response data:", string(data))
}
