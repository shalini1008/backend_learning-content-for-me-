package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websiteurl := []string{
		"https://amazon.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteurl {
		wg.Add(1)
		go printstatus(web)
	}

	wg.Wait()
}

func printstatus(web string) {
	defer wg.Done()
	res, err := http.Get(web)
	if err != nil {
		fmt.Println("oops issue in website url:", web)
		return
	}
	fmt.Printf("%d status of %s website\n", res.StatusCode, web)
}
