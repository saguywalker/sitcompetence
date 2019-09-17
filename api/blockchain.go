package api

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/cbergoon/merkletree"
	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GiveBadge takes a giving badge request and response with transactionID
func (a *API) GiveBadge(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	/*
		vals := r.URL.Query()
		params, ok := vals["data"]
		if !ok {
			return fmt.Errorf("missing data in parameter")
		}

		decoded, err := base64.StdEncoding.DecodeString(params[0])
		if err != nil {
			return err
		}
		ctx.Logger.Infof("decoded value\n%s\n", string(decoded))
	*/
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// Struct of objects
	var listOfBadges []*model.GiveBadge
	if err := json.Unmarshal(body, &listOfBadges); err != nil {
		return err
	}

	// Hash of each object will be used to make a merkle tree
	listOfHashes := make([]merkletree.Content, len(listOfBadges))
	for i, x := range listOfBadges {
		hash, err := x.CalculateHash()
		if err != nil {
			return err
		}

		listOfHashes[i] = model.MyHash(hash)
	}

	tree, err := merkletree.NewTree(listOfHashes)
	if err != nil {
		return err
	}

	ctx.Logger.Infof("Merkle Root Hash: %x", tree.MerkleRoot())
	/*
		hashes := make([]string, uint(2*math.Ceil(float64(len(listOfHashes))/2.0)))
		for i, leaf := range tree.Leafs {
			hashes[i] = fmt.Sprintf("%x", leaf.Hash)
		}
	*/
	transactionID, err := a.broadcastTX(ctx, w, tree.MerkleRoot())
	if err != nil {
		return err
	}

	if err := ctx.UpdateMerkleTransaction(transactionID, tree.MerkleRoot(), listOfHashes); err != nil {
		return err
	}

	if err := ctx.UpdateCollectedCompetence(listOfBadges, transactionID); err != nil {
		return err
	}

	return nil
}

// ApproveActivity takes an approving activity request and response with transactionID
func (a *API) ApproveActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var listOfActivities []model.ApproveActivity
	if err := json.Unmarshal(body, &listOfActivities); err != nil {
		return err
	}

	listOfHashes := make([]merkletree.Content, len(listOfActivities))
	for i, x := range listOfActivities {
		hash, err := x.CalculateHash()
		if err != nil {
			return err
		}

		listOfHashes[i] = model.MyHash(hash)
	}

	// Free listOfActivities
	listOfActivities = nil

	tree, err := merkletree.NewTree(listOfHashes)
	if err != nil {
		return err
	}

	hashes := make([]string, uint(2*math.Ceil(float64(len(listOfHashes))/2.0)))
	for i, leaf := range tree.Leafs {
		hashes[i] = fmt.Sprintf("%x", leaf.Hash)
	}

	transactionID, err := a.broadcastTX(ctx, w, tree.MerkleRoot())
	if err != nil {
		return nil
	}

	err = ctx.UpdateMerkleTransaction(transactionID, tree.MerkleRoot(), listOfHashes)

	return err
}

// VerifyTX verifies whethear a corresponding
func (a *API) VerifyTX(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var bodyMap map[string]interface{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		return nil
	}
	/*
		transactionID, err := hex.DecodeString(bodyMap["transaction_id"].(string))
		if err != nil {
			return nil
		}
	*/
	url := fmt.Sprintf("http://%s/tx?hash=0x%v", a.Config.Peers[a.CurrentPeerIndex], bodyMap["transaction_id"])
	ctx.Logger.Infoln(url)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	respData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var fullData map[string]interface{}
	err = json.Unmarshal(respData, &fullData)
	if err != nil {
		return err
	}

	respResult, ok := fullData["result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(fullData)
		return fmt.Errorf("error when asserting type from data[\"result\"]")
	}

	txResult, ok := respResult["tx_result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(respResult)
		return fmt.Errorf("error when asserting type from data[\"tx_result\"]")
	}

	events, ok := txResult["events"].([]interface{})
	if !ok {
		ctx.Logger.Errorln(txResult)
		return fmt.Errorf("error when asserting type from data[\"events\"]")
	}

	firstEvent, ok := events[0].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(events)
		return fmt.Errorf("error when asserting type from data[\"events[0]\"]")
	}

	attributes, ok := firstEvent["attributes"].([]interface{})
	if !ok {
		ctx.Logger.Errorln(firstEvent)
		return fmt.Errorf("error when asserting type from data[\"attributes\"]")
	}

	firstAttribute, ok := attributes[0].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(attributes)
		return fmt.Errorf("error when asserting type from data[\"attributes[0]\"]")
	}

	encodedValue := firstAttribute["value"]
	merkleRoot, err := base64.StdEncoding.DecodeString(encodedValue.(string))
	if err != nil {
		ctx.Logger.Errorf("%v", encodedValue)
		return err
	}
	ctx.Logger.Infof("merkle root: %x", merkleRoot)

	rawData, err := json.Marshal(bodyMap["data"])
	if err != nil {
		return err
	}
	ctx.Logger.Infof("json data\n%s\n", rawData)

	result, err := ctx.VerifyTX(merkleRoot, rawData)
	if err != nil {
		return err
	}
	ctx.Logger.Infoln(result)

	w.Write([]byte(fmt.Sprintf("%v\n", result)))

	return nil
}

// BroadcastTX broadcast transaction to blockchain node
// Note: Need to check wheater calling's node is reachable or not
func (a *API) broadcastTX(ctx *app.Context, w http.ResponseWriter, hash []byte) ([]byte, error) {
	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", a.Config.Peers[a.CurrentPeerIndex], hash)
	ctx.Logger.Infoln(url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var fullData map[string]interface{}
	err = json.Unmarshal(data, &fullData)
	if err != nil {
		return nil, err
	}

	resultData, ok := fullData["result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%+v\n", fullData)
		return nil, fmt.Errorf("error when asserting type from data[\"result\"]")
	}

	txIDInterface := resultData["hash"]
	transactionID, err := hex.DecodeString(txIDInterface.(string))
	if err != nil {
		return nil, err
	}
	ctx.Logger.Infof("TransactionID: %x\n", transactionID)

	//w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)

	// Move to the next peer in round-robin fashion
	a.CurrentPeerIndex = (a.CurrentPeerIndex + 1) % len(a.Config.Peers)

	return transactionID, err
}
