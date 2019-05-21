package main

import (
	"encoding/json"
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

	//	resp, _ := rpcClient.GetInfo()
	//	fmt.Println(resp)
	//	fmt.Println("****************************")
	//	resp, _ = rpcClient.ListStreams()
	//	fmt.Println(resp)
	//	fmt.Println("****************************")
	//	resp, _ = rpcClient.GetBlockchainParams()
	//	fmt.Println(resp)
	//	fmt.Println("****************************")
	//_, err = rpcClient.ListStreamItems("competencies")
	_, err = rpcClient.ListStreamQueryItems("competencies", "2018")
	if err != nil {
		panic(err)
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
