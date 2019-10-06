package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/model"
)

// GiveBadge hashing badge, broadcast it and update to database
func (ctx *Context) GiveBadge(badge *model.CollectedCompetence, w http.ResponseWriter, index uint64, peers []string) (uint64, error) {
	badgeHash, err := badge.CalculateHash()
	if err != nil {
		return index, err
	}

	txID, err := ctx.broadcastTX(badgeHash, index, peers)
	if err != nil {
		return index, err
	}

	badge.TxID = txID

	if err := ctx.Database.CreateCollectedCompetence(badge); err != nil {
		return index, err
	}

	index = (index + 1) % uint64(len(peers))

	return index, nil
}

// ApproveActivity hashing activity, broadcast it and update to database
func (ctx *Context) ApproveActivity(activity *model.AttendedActivity, w http.ResponseWriter, index uint64, peers []string) (uint64, error) {
	activityHash, err := activity.CalculateHash()
	if err != nil {
		return index, err
	}

	txID, err := ctx.broadcastTX(activityHash, index, peers)
	if err != nil {
		return index, err
	}

	activity.TransactionID = txID

	if err := ctx.Database.ApproveAttended(activity); err != nil {
		return index, err
	}

	index = (index + 1) % uint64(len(peers))

	return index, nil
}

func (ctx *Context) broadcastTX(hash []byte, index uint64, peers []string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", peers[index], hash)
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

	// Move to the next peer in round-robin fashion
	//ctx.CurrentPeerIndex = (ctx.CurrentPeerIndex + uint64(1)) % uint64(len(ctx.Peers))

	return transactionID, err
}

// VerifyTX verify data with a given merkle root
func (ctx *Context) VerifyTX(data []byte, index uint64, peers []string) (bool, uint64, []byte, error) {
	var trimData bytes.Buffer
	if err := json.Compact(&trimData, data); err != nil {
		return false, index, nil, err
	}
	hashData := sha256.Sum256(trimData.Bytes())
	ctx.Logger.Infof("H(data): %x\n", hashData)

	url := fmt.Sprintf("http://%s/abci_query?data=0x%x", peers[index], hashData)
	ctx.Logger.Infoln(url)

	index = (index + 1) % uint64(len(peers))

	resp, err := http.Get(url)
	if err != nil {
		return false, index, hashData[:], err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.Logger.Errorf("%v", resp)
		return false, index, hashData[:], err
	}

	var fullData map[string]interface{}
	if err := json.Unmarshal(respData, &data); err != nil {
		ctx.Logger.Errorf("%v", respData)
		return false, index, hashData[:], err
	}

	respResult, ok := fullData["result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%v", fullData)
		return false, index, hashData[:], err
	}

	respResult, ok = respResult["response"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%v", respResult)
		return false, index, hashData[:], err
	}

	isExist := respResult["log"] == "exists"

	return isExist, index, hashData[:], nil
}
