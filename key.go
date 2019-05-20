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

	rpcClient := jsonrpc.NewRPCClient(url)
	rpcClient.SetBasicAuth(config.Username, config.Password)

	/*jsonData := []byte(`{"params":["competencies",{"keys":["2","2018"]}]}`)
	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)

	mapData := make(map[string][]string)
	mapData["keys"] = []string{"2", "2018"}
	json2, _ := json.Marshal(mapData)
	var data2 [2]interface{}
	data2[0] = `competencies`
	data2[1] = `'` + string(json2) + `'`
	data2[1] = mapData
	fmt.Println(data2[1])
	fmt.Println("liststreamqueryitems", data2[0], data2[1])
	fmt.Println("****************************")*/
	tmp := [2]interface{}{StreamName, map[string][]string{"keys": []string{"59130500211"}}}
	listQueryCommand := Payload{
		Method: "liststreamqueryitems",
		Params: tmp,
	}
	fmt.Println("liststreamqueryitems", listQueryCommand.Params[0], listQueryCommand.Params[1])
	resp, err := rpcClient.Call("liststreamqueryitems", listQueryCommand.Params[0], listQueryCommand.Params[1])
	if err != nil {
		panic(err)
	}
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
