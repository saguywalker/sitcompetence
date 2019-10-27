package app

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/saguywalker/sitcompetence/model"
	dataTm "github.com/saguywalker/sitcompetence/proto/data"
	protoTm "github.com/saguywalker/sitcompetence/proto/tendermint"
	"golang.org/x/crypto/ed25519"
)

// GiveBadge hashing badge, broadcast it and update to database
func (ctx *Context) GiveBadge(badge *model.CollectedCompetence, sk string, index uint64, peers []string) (uint64, error) {
	giverPK, err := ctx.Database.GetStaffPublicKey(ctx.User.UserID)
	if err != nil {
		return index, err
	}

	badge.Giver = giverPK

	badgeProto := dataTm.CollectedBadge{
		StudentID: badge.StudentID,
		CompetenceID: badge.CompetenceID,
		Semester: badge.Semester,
		Giver: badge.Giver,
	}

	badgeBytes, err := proto.Marshal(&badgeProto)
	if err != nil {
		return index, err
	}

	txID, err := ctx.broadcastTX("GiveBadge", badgeBytes, giverPK, sk, index, peers)
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
func (ctx *Context) ApproveActivity(activity *model.AttendedActivity, index uint64, peers []string) (uint64, error) {
	approverPK, err := ctx.Database.GetStaffPublicKey(ctx.User.UserID)
	if err != nil {
		return index, err
	}

	activity.Approver = approverPK

	privKey := activity.PrivateKey
	activity.PrivateKey = ""

	approveBytes, err := json.Marshal(activity)
	if err != nil {
		return index, err
	}

	txID, err := ctx.broadcastTX("ApproveActivity", approveBytes, approverPK, privKey, index, peers)
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

func (ctx *Context) broadcastTX(method string, params, pubKey []byte, privKey string, index uint64, peers []string) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	// Decode base64 private key
	priv, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return nil, err
	}

	// Sign
	signature := ed25519.Sign(priv, params)

	payload := protoTm.Payload{
		Method: method,
		Params: string(params),
	}

	tx := protoTm.Tx{
		Payload:   &payload,
		PublicKey: pubKey,
		Signature: signature,
	}

	txBytes, err := proto.Marshal(&tx)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	for i := 0; i < len(peers); i++ {
		url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=%v", peers[index], txBytes)
		ctx.Logger.Infoln(url)
		resp, err := httpClient.Get(url)
		index = uint64((index + 1)) % uint64(len(peers))
		if err != nil {
			continue
		} else {
			data, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				continue
			}
			break
		}
	}

	var fullData map[string]interface{}
	if err := json.Unmarshal(data, &fullData); err != nil {
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

	return transactionID, err
}

// VerifyTX verify data with a given merkle root
func (ctx *Context) VerifyTX(data []byte, index uint64, peers []string) (bool, uint64, error) {
	var trimData bytes.Buffer
	if err := json.Compact(&trimData, data); err != nil {
		return false, index, err
	}
	ctx.Logger.Infof("data: %v\n", trimData.Bytes())

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	respData := make([]byte, 0)
	for i := 0; i < len(peers); i++ {
		url := fmt.Sprintf("http://%s/abci_query?data=%v", peers[index], trimData.Bytes())
		ctx.Logger.Infoln(url)
		resp, err := httpClient.Get(url)
		index = uint64((index + 1)) % uint64(len(peers))
		if err != nil {
			continue
		} else {
			respData, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				continue
			}
			ctx.Logger.Infoln(string(respData))
			break
		}
	}

	var fullData map[string]interface{}
	if err := json.Unmarshal(respData, &fullData); err != nil {
		ctx.Logger.Errorf("%v", respData)
		return false, index, err
	}

	respResult, ok := fullData["result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%v", fullData)
		return false, index, errors.New(string(respData))
	}

	respResult, ok = respResult["response"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%v", respResult)
		return false, index, errors.New(string(respData))
	}

	isExist := respResult["log"] == "exists"

	return isExist, index, nil
}

// GetBadgeFromStudent return all of collected badges from corresponding studentId from in blockchain
func (ctx *Context) GetBadgeFromStudent(id string, index uint64, peers []string) ([]model.CollectedCompetence, error) {
	return nil, nil
}
