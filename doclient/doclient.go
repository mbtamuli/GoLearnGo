package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	baseurl := "https://api.digitalocean.com/"
	doToken := os.Getenv("DO_TOKEN")
	if len(doToken) == 0 {
		fmt.Println("Run 'export DO_TOKEN=mySecretDigitalOceanToken'")
		os.Exit(1)
	}
	bearer := "Bearer " + doToken

	/*
	var jsonStr = []byte(`{
		// "name": "one"
		"names": [
			"two",
			"three"
		],
		"region": "blr1",
		"size": "s-1vcpu-1gb",
		"image": "ubuntu-18-04-x64",
		"ssh_keys": null,
		"backups": false,
		"ipv6": true,
		"user_data": null,
		"private_networking": null,
		"volumes": null,
		"tags": [
			"mbtamuli"
		]
	}`)

	req, err := http.NewRequest(http.MethodPost, baseurl+"v2/droplets", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string([]byte(body)))
	*/


	type Droplet struct {
		id int
		name string
	}

	req, err := http.NewRequest(http.MethodGet, baseurl+"v2/droplets", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var d Droplet
	err := json.Unmarshal([]byte(body), &d)
	fmt.Println(d.name)


	// Delete droplet /v2/droplets/$DROPLET_ID
	// Delete droplet with tag /v2/droplets?tag_name=$TAG_NAME
	/*
	req, err := http.NewRequest(http.MethodDelete, baseurl+"v2/droplets?tag_name=mbtamuli", nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode == 204 {
		fmt.Printf("Successfully deleted droplets with tag mbtamuli")
	} else {
		panic("Failure")
	}
	*/
}
