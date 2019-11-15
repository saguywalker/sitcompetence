package app

import (
	"bytes"
	"crypto/sha256"
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
	protoTm "github.com/saguywalker/sitcompetence/proto/tendermint"
	"golang.org/x/crypto/ed25519"
)

// GiveBadge hashing badge, broadcast it and update to database
func (ctx *Context) GiveBadge(badge *model.CollectedCompetence, index uint64, peers []string, webSK []byte) ([]byte, uint64, error) {
	ctx.Logger.Infof("app/GiveBadge: %v\n", *badge)

	badgeBytes, err := json.Marshal(badge)
	if err != nil {
		return nil, index, err
	}

	txID, err := ctx.broadcastTX("GiveBadge", badgeBytes, index, peers, webSK)
	if err != nil {
		return txID, index, err
	}

	index = (index + 1) % uint64(len(peers))

	return txID, index, nil
}

// ApproveActivity hashing activity, broadcast it and update to database
func (ctx *Context) ApproveActivity(activity *model.AttendedActivity, index uint64, peers []string, webSK []byte) ([]byte, uint64, error) {
	approveBytes, err := json.Marshal(activity)
	if err != nil {
		return nil, index, err
	}

	txID, err := ctx.broadcastTX("ApproveActivity", approveBytes, index, peers, webSK)
	if err != nil {
		return txID, index, err
	}

	if err := ctx.Database.ApproveAttended(activity); err != nil {
		return txID, index, err
	}

	index = (index + 1) % uint64(len(peers))

	return txID, index, nil
}

func (ctx *Context) broadcastTX(method string, params []byte, index uint64, peers []string, webSK []byte) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	hashParams := sha256.Sum256(params)
	tmSignature := ed25519.Sign(webSK, hashParams[:])

	payload := protoTm.Payload{
		Method: method,
		Params: params,
	}

	tx := protoTm.Tx{
		Payload:   &payload,
		Signature: tmSignature,
	}

	txBytes, err := proto.Marshal(&tx)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	for i := 0; i < len(peers); i++ {
		url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", peers[index], txBytes)
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
func (ctx *Context) VerifyTX(data []byte, index uint64, peers []string) (bool, []byte, uint64, error) {
	var trimData bytes.Buffer
	if err := json.Compact(&trimData, data); err != nil {
		return false, nil, index, err
	}
	ctx.Logger.Infof("data: %s\n", trimData.String())

	_, evidence, returnIndex, err := ctx.BlockchainQueryWithParams(trimData.String(), index, peers)
	if err != nil {
		if err.Error() == "does not exists" {
			return false, evidence, returnIndex, nil
		}

		return false, evidence, returnIndex, err
	}

	return true, evidence, returnIndex, nil
}

// GetBadgeFromStudent return all of collected badges from corresponding studentId from in blockchain
func (ctx *Context) GetBadgeFromStudent(id string, index uint64, peers []string) ([]model.CollectedCompetence, error) {
	return nil, nil
}

// BlockchainQueryWithParams return []byte from corresponding parameters
func (ctx *Context) BlockchainQueryWithParams(params string, index uint64, peers []string) ([]byte, []byte, uint64, error) {
	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	respData := make([]byte, 0)
	for i := 0; i < len(peers); i++ {
		url := fmt.Sprintf("http://%s/abci_query?data=0x%x", peers[index], []byte(params))
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
		ctx.Logger.Errorf("%s", respData)
		return nil, respData, index, err
	}

	respResult, ok := fullData["result"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%s", fullData)
		return nil, respData, index, errors.New(string(respData))
	}

	respResult, ok = respResult["response"].(map[string]interface{})
	if !ok {
		ctx.Logger.Errorf("%s", respResult)
		return nil, respData, index, errors.New(string(respData))
	}

	if isExist := respResult["log"] == "exists"; !isExist {
		return nil, respData, index, errors.New("does not exists")
	}

	dec64, err := base64.StdEncoding.DecodeString(respResult["value"].(string))
	if err != nil {
		return nil, respData, index, err
	}

	return dec64, respData, index, nil
}

// VerifySignature authenticate the user and its signature
func (ctx *Context) VerifySignature(message []byte, signature, b64PubKey string) (bool, error) {
	hashed := sha256.Sum256(message)

	pubKey, err := base64.StdEncoding.DecodeString(b64PubKey)
	if err != nil {
		ctx.Logger.Errorln("error when decoding publickey")
		return false, err
	}

	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		ctx.Logger.Errorln("error when decoding signature")
		return false, err
	}

	isVerified := ed25519.Verify(pubKey, hashed[:], sig)

	ctx.Logger.Infof("params: %s\nsha256: %x\nsignature: %x\nisVerified: %v\n", message, hashed[:], sig, isVerified)

	return isVerified, nil
}
