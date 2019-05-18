package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config AccountConfig
	json.Unmarshal(byteValue, &config)
	fmt.Println(config)

	listStreams := Payload{
		ID:     0,
		Method: "liststreams",
	}

	listStreamsBytes, err := json.Marshal(listStreams)
	if err != nil {
		panic(err)
	}

	url := "http://" + config.IP + ":" + config.Port

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(listStreamsBytes))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(config.Username, config.Password)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

type AccountConfig struct {
	Username string
	Password string
	IP       string
	Port     string
	Chain    string
}

type Payload struct {
	ID     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}
