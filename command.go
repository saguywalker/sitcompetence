package main

import (
	"fmt"

	"github.com/KeisukeYamashita/jsonrpc"
)

type Client struct {
	jsonrpc.RPCClient
}

func (client *Client) GetInfo() (*jsonrpc.RPCResponse, error) {
	getInfoCommand := Payload{
		Method: "getinfo",
	}

	resp, err := client.Call(getInfoCommand.Method)

	return resp, err
}

func (client *Client) ListStreams() (*jsonrpc.RPCResponse, error) {
	listStreamsCommand := Payload{
		Method: "liststreams",
	}

	resp, err := client.Call(listStreamsCommand.Method)

	return resp, err
}

func (client *Client) ListStreamQueryItems(stream string, keys ...string) ([]CollectedCompetencies, error) {
	params := []interface{}{stream, map[string][]string{"keys": keys}}
	listStreamQueryItemsCommand := Payload{
		Method: "liststreamqueryitems",
		Params: params,
	}
	resp, err := client.Call(listStreamQueryItemsCommand.Method, listStreamQueryItemsCommand.Params[0], listStreamQueryItemsCommand.Params[1])
	if err != nil {
		return nil, err
	}

	var data []interface{}
	resp.GetObject(&data)
	output := make([]CollectedCompetencies, len(data))

	for i, x := range data {
		x2 := x.(map[string]interface{})
		output[i] = CollectedCompetencies{
			x2["txid"].(string),
			x2["keys"].([]interface{}),
			x2["data"].(map[string]interface{}),
			x2["blocktime"].(float64),
			x2["publishers"].([]interface{}),
		}
	}

	return output, err
}

func (client *Client) GetBlockchainParams() (*jsonrpc.RPCResponse, error) {
	getBlockchainParamsCommand := Payload{
		Method: "getblockchainparams",
	}

	resp, err := client.Call(getBlockchainParamsCommand.Method)

	return resp, err
}

func (client *Client) ListStreamItems(stream string) ([]CollectedCompetencies, error) {
	params := []interface{}{stream}
	listStreamItemsCommand := Payload{
		Method: "liststreamitems",
		Params: params,
	}

	resp, err := client.Call(listStreamItemsCommand.Method, listStreamItemsCommand.Params[0])

	var data []interface{}
	resp.GetObject(&data)
	output := make([]CollectedCompetencies, 0)

	for _, x := range data {
		x2 := x.(map[string]interface{})
		tmp, ok := x2["data"].(map[string]interface{})
		if !ok {
			continue
		}
		tmp2, ok := tmp["json"].(map[string]interface{})
		if !ok {
			continue
		}

		tmpOutput := CollectedCompetencies{
			x2["txid"].(string),
			x2["keys"].([]interface{}),
			tmp2,
			x2["blocktime"].(float64),
			x2["publishers"].([]interface{}),
		}

		output = append(output, tmpOutput)
	}

	return output, err
}

func (client *Client) PublishManually(stream string, keys []string, studentId string, competenceId string, staffId string) (*jsonrpc.RPCResponse, error) {
	mapData := map[string]string{"student_id": studentId, "competence_id": competenceId, "staff_id": staffId}

	/*jsonData, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}*/
	mapJson := map[string]interface{}{"json": mapData}

	params := []interface{}{stream, keys, mapJson}
	publishCommand := Payload{
		Method: "publish",
		Params: params,
	}
	fmt.Println(publishCommand.Method, publishCommand.Params)
	resp, err := client.Call(publishCommand.Method, publishCommand.Params[0], publishCommand.Params[1], publishCommand.Params[2])
	if err != nil {
		panic(err)
	}

	return resp, err
}

type CollectedCompetencies struct {
	Txid       string
	Keys       []interface{}
	Data       map[string]interface{}
	Blocktime  float64
	Publishers []interface{}
}
