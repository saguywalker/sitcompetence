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

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config AccountConfig
	json.Unmarshal(byteValue, &config)
	fmt.Println(config)

	/*listStreamsBytes, err := json.Marshal(Payload{
		Method: "liststreams",
	})
	if err != nil {
		panic(err)
	}
	listStreams := Payload{
		Method: "liststreams",
	}*/
	//fmt.Println(listStreams)

	url := "http://" + config.IP + ":" + config.Port

	rpcClient := jsonrpc.NewRPCClient(url)
	rpcClient.SetBasicAuth(config.Username, config.Password)

	//jsonData := []byte(`{"params":["competencies",{"keys":["2","2018"]}]}`)
	//jsonData := []byte(`{"keys":["2","2018"]}`)
	jsonData := []byte(`{"method":"liststreamqueryitems","params":["competencies",{"keys":["2","2018"]}]}`)
	var data Payload
	json.Unmarshal(jsonData, &data)
	fmt.Println(string(jsonData))
	fmt.Println(data)
	//data := v.(map[string]interface{})
	/*
		for k, v := range data {
			switch v := v.(type) {
			case string:
				fmt.Println(k, v, "(string)")
			case float64:
				fmt.Println(k, v, "(float)")
			case []interface{}:
				fmt.Println(k, v, "(array)")
				for i, u := range v {
					fmt.Println("	", i, u)
				}
			default:
				fmt.Println(k, v, "(unknown)")
			}
		}*/
	/*
		resp, err := rpcClient.Call(listStreamQueryItems.Method, listStreamQueryItems.Params)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)*/
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
	Method string `json:"method"`
	Params struct {
		StreamName string                 `json:"stream"`
		Keys       map[string]interface{} `json:"keys"`
	} `json:"params"`
}
