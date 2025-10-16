package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Golang Bootcamp", 599, "LearnCodeOnline.in", "abxd123", []string{"backedn", "Go"}},
		{"Angular Bootcamp", 799, "LearnCodeOnline.in", "gffxd123", []string{}},
	}

	//package this data as Json data

	finalJson, _ := json.MarshalIndent(lcoCourse, "", "\t")
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`[
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": [
				"web-dev",
				"js"
			]
		},
		{
			"coursename": "Golang Bootcamp",
			"Price": 599,
			"website": "LearnCodeOnline.in",
			"tags": [
				"backend",
				"Go"
			]
		}
	]`)

	var lcoCourses []course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("json was not valid")
	}

	var myOnlineData []map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T", k, v, v)
	}
}
