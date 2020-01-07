package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	type Slides struct {
		Title string   `json:"title"`
		Type  string   `json:"type"`
		Items []string `json:"items,omitempty"`
	}
	type Slideshow struct {
		Author string   `json:"author"`
		Date   string   `json:"date"`
		Slides []Slides `json:"slides"`
		Title  string   `json:"title"`
	}
	type Container struct {
		Slideshow Slideshow `json:"slideshow"`
	}

	var obj Container

	jsonFile, err := os.Open("slides.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Json Object: \n%+v\n\n", obj)

	for i := 0; i < len(obj.Slideshow.Slides); i++ {
		fmt.Println("Titles: " + obj.Slideshow.Slides[i].Title)
	}
}
