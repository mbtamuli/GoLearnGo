package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	type Animal struct {
		Name  string `json:"Name"`
		Order string `json:"Order"`
	}
	var animals []Animal

	jsonFile, err := os.Open("animals.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &animals)
	for i := 0; i < len(animals); i++ {
		fmt.Println("Name: " + animals[i].Name)
		fmt.Println("Order: " + animals[i].Order)
	}

}
