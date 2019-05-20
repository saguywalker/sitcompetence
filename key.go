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

	url := "http://" + config.IP + ":" + config.Port

	rpcClient := jsonrpc.NewRPCClient(url)
	rpcClient.SetBasicAuth(config.Username, config.Password)

	jsonData := []byte(`{"params":["competencies",{"keys":["2","2018"]}]}`)
	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)
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
				switch u := u.(type) {
				case map[string]interface{}:
					for kk, vv := range u {
						fmt.Println("		", kk, vv, "(sub map)")
						vvv := vv.([]interface{})
						for j, jj := range vvv {
							fmt.Println("			", j, jj)
						}
					}
				case []interface{}:
					for ii, vv := range u {
						fmt.Println("		", ii, vv, "(sub array)")
					}
				default:
					fmt.Println("		", v, "(sub unknown)")
				}
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
	fmt.Println("***********************************")

	mapData := make(map[string][]string)
	mapData["keys"] = []string{"2", "2018"}
	json2, _ := json.Marshal(mapData)
	var data2 [2]interface{}
	data2[0] = "competencies"
	data2[1] = `'` + string(json2) + `'`
	fmt.Println(data2[1])
	fmt.Println("liststreamqueryitems", data2[0], data2[1])
	resp, err := rpcClient.Call("liststreamqueryitems", data2)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

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
