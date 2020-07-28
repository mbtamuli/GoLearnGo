package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/jbowtie/gokogiri"
)

func main() {
	inputFile := flag.String("input", "./System.xml", "Input XML file. Defaults to './System.xml'")
	propertyAttribute := flag.String("property", "", "Name of Property to edit. Example: ENABLE_PLAIN_CSV")
	value := flag.String("value", "", "Value that needs to be set for the Property. Example: YES")
	flag.Parse()

	xmlSource, err := ioutil.ReadFile(*inputFile)

	if err != nil {
		log.Fatalln(err)
	}

	doc, _ := gokogiri.ParseXml(xmlSource)

	idNode, _ := doc.Search("/MODULE/SYSTEM/PROPERTIES/PROPERTY[@NAME='" + *propertyAttribute + "']")
	if len(idNode) > 0 {
		idNode[0].SetContent(*value)
	}

	str2, _ := doc.ToXml(nil, nil)
	err = ioutil.WriteFile("./Modified.xml", str2, 0644)

	if err != nil {
		log.Fatalln(err)
	}
}
