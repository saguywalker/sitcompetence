package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/KeisukeYamashita/jsonrpc"
)

const (
	//StreamName for testing.
	StreamName string = "competencies"
)

func main() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config AccountConfig
	json.Unmarshal(byteValue, &config)

	url := "http://" + config.IP + ":" + config.Port

	client := jsonrpc.NewRPCClient(url)
	rpcClient := Client{
		*client,
	}
	rpcClient.SetBasicAuth(config.Username, config.Password)

	/*a, err := rpcClient.PublishManually("competencies", []string{"59130500218", "1", "2018"}, "59130500218", "CollabMindset_1", "12344321")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)*/
	/*
		collectedCompetencies, err := rpcClient.ListStreamItems("competencies")
		if err != nil {
			panic(err)
		}

		for i, x := range collectedCompetencies {
			fmt.Println(i)
			fmt.Println("txid: ", x.Txid)
			fmt.Print("keys:")
			for _, k := range x.Keys {
				fmt.Print(k, " ")
			}
			fmt.Print("\ndata: ")
			for k, v := range x.Data {
				fmt.Print("(", k, ":", v, ") ")
			}
			fmt.Println("\nblocktime: ", x.Blocktime)
			fmt.Println("publisher(s): ", x.Publishers[0])
			fmt.Println("************************************")
		}*/

	lastBlock, err := rpcClient.GetLastBlockInfo()
	if err != nil {
		panic(err)
	}
	lastBlockMap := lastBlock.(map[string]interface{})
	for k, v := range lastBlockMap {
		fmt.Println("key:", k, ", value:", v)
	}
	fmt.Println("************************************")

	block24, err := rpcClient.GetBlock(24)
	if err != nil {
		panic(err)
	}
	block24Map := block24.(map[string]interface{})
	for k, v := range block24Map {
		fmt.Println("key:", k, ", value:", v)
	}

}

//AccountConfig uses for connect to multichain
type AccountConfig struct {
	Username string
	Password string
	IP       string
	Port     string
	Chain    string
}
