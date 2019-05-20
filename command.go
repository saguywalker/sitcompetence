package main

import (
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

func (client *Client) ListStreamQueryItems(stream string, keys ...string) (*jsonrpc.RPCResponse, error) {
	params := [2]interface{}{stream, map[string][]string{"keys": keys}}
	listStreamQueryItemsCommand := Payload{
		Method: "liststreamqueryitems",
		Params: params,
	}

	resp, err := client.Call(listStreamQueryItemsCommand.Method, listStreamQueryItemsCommand.Params[0], listStreamQueryItemsCommand.Params[1])

	return resp, err
}

func (client *Client) GetBlockchainParams() (*jsonrpc.RPCResponse, error) {
	getBlockchainParamsCommand := Payload{
		Method: "getblockchainparams",
	}

	resp, err := client.Call(getBlockchainParamsCommand.Method)

	return resp, err
}
