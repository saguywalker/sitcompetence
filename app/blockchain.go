package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/model"
)

// GiveBadge hashing badge, broadcast it and update to database
func (ctx *Context) GiveBadge(badge *model.CollectedCompetence, w http.ResponseWriter, peers []string) error {
	badgeHash, err := badge.CalculateHash()
	if err != nil {
		return err
	}

	txID, err := ctx.broadcastTX(badgeHash, peers)
	if err != nil {
		return err
	}

	badge.TxID = txID

	if err := ctx.Database.CreateCollectedCompetence(badge); err != nil {
		return err
	}

	return nil
}

// ApproveActivity hashing activity, broadcast it and update to database
func (ctx *Context) ApproveActivity(activity *model.AttendedActivity, w http.ResponseWriter, peers []string) error {
	activityHash, err := activity.CalculateHash()
	if err != nil {
		return err
	}

	txID, err := ctx.broadcastTX(activityHash, peers)
	if err != nil {
		return err
	}

	activity.TransactionID = txID

	if err := ctx.Database.ApproveAttended(activity); err != nil {
		return err
	}

	return nil
}

func (ctx *Context) broadcastTX(hash []byte, peers []string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", peers[ctx.CurrentPeerIndex], hash)
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
	ctx.CurrentPeerIndex = (ctx.CurrentPeerIndex + uint64(1)) % uint64(len(ctx.Peers))

	return transactionID, err
}

// VerifyTX verify data with a given merkle root
func (ctx *Context) VerifyTX(data, txID []byte) (bool, error) {
	var trimData bytes.Buffer
	if err := json.Compact(&trimData, data); err != nil {
		return false, err
	}
	hashData := sha256.Sum256(trimData.Bytes())
	ctx.Logger.Infof("H(data): %x\n", hashData)

	url := fmt.Sprintf("http://%s/tx?hash=0x%x", ctx.Peers[ctx.CurrentPeerIndex], txID)
	ctx.CurrentPeerIndex = (ctx.CurrentPeerIndex + uint64(1)) % uint64(len(ctx.Peers))
	ctx.Logger.Infoln(url)

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	respData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	var fullData map[string]interface{}
	err = json.Unmarshal(respData, &fullData)
	if err != nil {
		return false, err
	}

	respResult, ok := fullData["result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(fullData)
		return false, fmt.Errorf("error when asserting type from data[\"result\"]")
	}

	txResult, ok := respResult["tx_result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(respResult)
		return false, fmt.Errorf("error when asserting type from data[\"tx_result\"]")
	}

	events, ok := txResult["events"].([]interface{})
	if !ok {
		ctx.Logger.Errorln(txResult)
		return false, fmt.Errorf("error when asserting type from data[\"events\"]")
	}

	firstEvent, ok := events[0].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(events)
		return false, fmt.Errorf("error when asserting type from data[\"events[0]\"]")
	}

	attributes, ok := firstEvent["attributes"].([]interface{})
	if !ok {
		ctx.Logger.Errorln(firstEvent)
		return false, fmt.Errorf("error when asserting type from data[\"attributes\"]")
	}

	firstAttribute, ok := attributes[0].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorln(attributes)
		return false, fmt.Errorf("error when asserting type from data[\"attributes[0]\"]")
	}

	encodedValue := firstAttribute["value"]
	blockValue, err := base64.StdEncoding.DecodeString(encodedValue.(string))
	if err != nil {
		ctx.Logger.Errorf("%v", encodedValue)
		return false, err
	}
	ctx.Logger.Infof("data on blockchain: %x", blockValue)

	if bytes.Equal(blockValue, hashData[:]) {
		return true, nil
	}

	return false, nil
}
