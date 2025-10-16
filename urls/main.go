package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=12345"

func main() {
	fmt.Println("Welcome to url in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Path:   "www.nextgenresearch.co.in",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
