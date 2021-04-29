package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mbtamuli/GoLearnGo/protobufs/pb"
	"gopkg.in/yaml.v3"
)

func createMessage() {
	data := &pb.Person{
		Name:  "Harry Potter",
		Id:    1,
		Email: "harry@potter.com",
	}
	fmt.Println(data)
}

func createMessageFromJSON() error {
	bytes, err := ioutil.ReadFile("testdata/person1.json")
	if err != nil {
		return err
	}

	data := &pb.Person{}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	fmt.Println(data)
	return nil
}

func createMessageFromYAML() error {
	bytes, err := ioutil.ReadFile("testdata/person1.yaml")
	if err != nil {
		return err
	}

	data := &pb.Person{}
	err = yaml.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	fmt.Println(data)
	return nil
}

func main() {
	createMessage()
	createMessageFromJSON()
	createMessageFromYAML()
}
