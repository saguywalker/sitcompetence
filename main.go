package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/KeisukeYamashita/jsonrpc"
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

	/*listStreamsBytes, err := json.Marshal(Payload{
		Method: "liststreams",
	})
	if err != nil {
		panic(err)
	}*/
	listStreams := Payload{
		Method: "liststreams",
	}

	url := "http://" + config.IP + ":" + config.Port

	rpcClient := jsonrpc.NewRPCClient(url)
	rpcClient.SetBasicAuth(config.Username, config.Password)
	fmt.Println(rpcClient)

	resp, err := rpcClient.Call("getinfo")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println("******************************")

	resp, err = rpcClient.Call(listStreams.Method, listStreams.Params)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	/*
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
	*/

}

type AccountConfig struct {
	Username string
	Password string
	IP       string
	Port     string
	Chain    string
}

type Payload struct {
	Method string
	Params []interface{}
}
