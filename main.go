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
	rpcClient := Client{
		*client,
	}
	rpcClient.SetBasicAuth(config.Username, config.Password)

	/*a, err := rpcClient.PublishManually("competencies", []string{"59130500211", "1", "2019"}, "59130500211", "growthmindset1", "12345678")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)*/
	fmt.Println("************************************")
	//a, err = rpcClient.PublishManually("competencies", []string{"59130500210", "2", "2019"}, "59130500210", "communicating1", "12312312")
	//a, err = rpcClient.PublishManually("competencies", []string{"59130500218", "1", "2018"}, "59130500218", "globalmindset1", "12341234")

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
	}
	/*var data []interface{}
	err = resp.GetObject(&data)
	data2 := data[0].(map[string]interface{})
	keys := make([]string, 0, len(data2))

	for k := range data2 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)

	for i, x := range data {
		x2 := x.(map[string]interface{})
		j := 0
		for j < len(x2) {
			fmt.Println(i, keys[j], x2[keys[j]], reflect.TypeOf(x2[keys[j]]))
			j++
		}
		fmt.Println("********************************")
	}*/

}

type AccountConfig struct {
	Username string
	Password string
	IP       string
	Port     string
	Chain    string
}

type Payload struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}
