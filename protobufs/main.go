package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/mbtamuli/GoLearnGo/protobufs/pb"
	"google.golang.org/protobuf/encoding/protojson"
)

func createMessageFromYAML() (*pb.Person, error) {
	yamlBytes, err := ioutil.ReadFile("testdata/person.yaml")
	if err != nil {
		return nil, err
	}

	person := &pb.Person{}
	err = yaml.Unmarshal(yamlBytes, person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func createMessageFromYAMLWithJSON() (*pb.Person, error) {
	yamlBytes, err := ioutil.ReadFile("testdata/person.yaml")
	if err != nil {
		return nil, err
	}

	person := &pb.Person{}
	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(jsonBytes, person); err != nil {
		return nil, err
	}

	return person, nil
}

func main() {
	p1, err := createMessageFromYAML()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Image URL when using just yaml.Unmarshal: " + p1.GetImageUrl())

	p2, err := createMessageFromYAMLWithJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Image URL when using YAMLToJSON then protojson.Unmarshal: " + p2.GetImageUrl())

}
