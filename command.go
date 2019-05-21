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
	fmt.Println(params)
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

func (client *Client) PublishManually(stream string, keys []string, studentId string, competenceId string, staffId string) (*jsonrpc.RPCResponse, error) {
	mapData := map[string]string{"student_id": studentId, "competence_id": competenceId, "staff_id": staffId}
	params := []interface{}{stream, keys, mapData}
	fmt.Println(params)
	publishCommand := Payload{
		Method: "publish",
		Params: params,
	}
	resp, err := client.Call(publishCommand.Method, publishCommand.Params)

	return resp, err
}

type CollectedCompetencies struct {
	Txid       string
	Keys       []interface{}
	Data       map[string]interface{}
	Blocktime  float64
	Publishers []interface{}
}
