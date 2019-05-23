package main

import (
	"github.com/KeisukeYamashita/jsonrpc"
)

//Client embed jsonrpc.RPCClient.
type Client struct {
	jsonrpc.RPCClient
}

//GetInfo calls getinfo command.
func (client *Client) GetInfo() (*jsonrpc.RPCResponse, error) {
	getInfoCommand := Payload{
		Method: "getinfo",
	}

	resp, err := client.Call(getInfoCommand.Method)

	return resp, err
}

//ListStreams calls liststreams command.
func (client *Client) ListStreams() ([]map[string]interface{}, error) {
	listStreamsCommand := Payload{
		Method: "liststreams",
	}

	resp, err := client.Call(listStreamsCommand.Method)
	var tmpResp interface{}
	resp.GetObject(&tmpResp)
	tmp := tmpResp.([]interface{})
	output := make([]map[string]interface{}, len(tmp))
	for i, x := range tmp {
		x2 := x.(map[string]interface{})
		output[i] = x2
	}
	return output, err
}

//ListStreamQueryItems calls liststreamqueryitems with stream name and key(s).
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

//GetBlockchainParams calls getblockchainparams.
func (client *Client) GetBlockchainParams() (*jsonrpc.RPCResponse, error) {
	getBlockchainParamsCommand := Payload{
		Method: "getblockchainparams",
	}

	resp, err := client.Call(getBlockchainParamsCommand.Method)

	return resp, err
}

//ListStreamItems calls listStreamItems from stream name without any key.
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
		tmpJSON, ok := tmp["json"].(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := x2["blocktime"].(float64); !ok {
			continue
		}

		tmpOutput := CollectedCompetencies{
			x2["txid"].(string),
			x2["keys"].([]interface{}),
			tmpJSON,
			x2["blocktime"].(float64),
			x2["publishers"].([]interface{}),
		}

		output = append(output, tmpOutput)
	}

	return output, err
}

//PublishManually is used for give a competence id manually by staff.
func (client *Client) PublishManually(stream string, keys []string, studentID string, competenceID string, staffID string) (*jsonrpc.RPCResponse, error) {
	mapData := map[string]string{"student_id": studentID, "competence_id": competenceID, "staff_id": staffID}

	mapJSON := map[string]interface{}{"json": mapData}

	params := []interface{}{stream, keys, mapJSON}
	publishCommand := Payload{
		Method: "publish",
		Params: params,
	}
	resp, err := client.Call(publishCommand.Method, publishCommand.Params[0], publishCommand.Params[1], publishCommand.Params[2])
	if err != nil {
		panic(err)
	}

	return resp, err
}

//PublishWithActivity is used for adding competence according to an activity.
func (client *Client) PublishWithActivity(stream string, keys []string, studentID string, competenceID string, activityID string) (*jsonrpc.RPCResponse, error) {
	mapData := map[string]string{"student_id": studentID, "competence_id": competenceID, "staff_id": activityID}

	mapJSON := map[string]interface{}{"json": mapData}

	params := []interface{}{stream, keys, mapJSON}
	publishCommand := Payload{
		Method: "publish",
		Params: params,
	}
	resp, err := client.Call(publishCommand.Method, publishCommand.Params[0], publishCommand.Params[1], publishCommand.Params[2])
	if err != nil {
		panic(err)
	}

	return resp, err
}

//GetLastBlockInfo return the last block's information
func (client *Client) GetLastBlockInfo() (interface{}, error) {
	getInfoCommand := Payload{
		Method: "getlastblockinfo",
	}

	resp, err := client.Call(getInfoCommand.Method)
	var output interface{}
	resp.GetObject(&output)

	return output, err
}

//GetBlock return the corresponding block number
func (client *Client) GetBlock(number uint32) (interface{}, error) {
	params := []interface{}{number}
	getInfoCommand := Payload{
		Method: "getblock",
		Params: params,
	}

	resp, err := client.Call(getInfoCommand.Method, getInfoCommand.Params[0])
	var output interface{}
	resp.GetObject(&output)

	return output, err
}

//CollectedCompetencies for unmarshal json from Blockchain.
type CollectedCompetencies struct {
	Txid       string
	Keys       []interface{}
	Data       map[string]interface{}
	Blocktime  float64
	Publishers []interface{}
}

//Payload uses to pass the arguments to multichain.
type Payload struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}
