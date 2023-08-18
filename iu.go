package main

import (
	"encoding/json"
	"fmt"
)

type couressr struct {
	Name       string `json:"Name"`
	Level      string `json:"Level"`
	View       int    `json:"View"`
	Instructor string `json:"Instructor"`
	FullPrice  int    `json:"FullPrice"`
}

func main() {
	data := []byte(`{
		Name:       "bc go",
		Level:      "nor",
		View:       7919,
		Instructor: "zen",
		FullPrice: 9999,
	}`)

	var cc couressr
	err := json.Unmarshal(data, &cc)
	fmt.Printf("%v\n", cc)
	// b, err := json.Marshal(c)
	// fmt.Printf("Type : %T \n ", b)
	// fmt.Printf("Byte : %v \n", b)
	// fmt.Printf("String : %s \n", b)

	fmt.Println(err)
}
