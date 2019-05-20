package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/KeisukeYamashita/jsonrpc"
)

const (
	StreamName string = "competencies"
)

func main() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config AccountConfig
	json.Unmarshal(byteValue, &config)

	url := "http://" + config.IP + ":" + config.Port

	client := jsonrpc.NewRPCClient(url)
	fmt.Println(client)
	rpcClient := Client{
		*client,
	}
	fmt.Println(rpcClient)
	rpcClient.SetBasicAuth(config.Username, config.Password)

	resp, _ := rpcClient.GetInfo()
	fmt.Println(resp)
	fmt.Println("****************************")
	resp, _ = rpcClient.ListStreams()
	fmt.Println(resp)
	fmt.Println("****************************")
	resp, _ = rpcClient.GetBlockchainParams()
	fmt.Println(resp)
	fmt.Println("****************************")
	resp, _ = rpcClient.ListStreamQueryItems("competencies", "2", "2018")
	fmt.Println(resp)
	fmt.Println("****************************")

}

type AccountConfig struct {
	Username string
	Password string
	IP       string
	Port     string
	Chain    string
}

type Payload struct {
	Method string         `json:"method"`
	Params [2]interface{} `json:"params"`
}
