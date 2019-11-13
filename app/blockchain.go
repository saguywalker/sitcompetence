package app

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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
func (ctx *Context) GiveBadge(badge *model.CollectedCompetence, sk string, index uint64, peers []string, key []byte) ([]byte, uint64, error) {
	ctx.Logger.Infof("app/GiveBadge: %v, %s\n", *badge, sk)

	giverPK, err := ctx.Database.GetStaffPublicKey(ctx.User.UserID)
	if err != nil {
		return nil, index, err
	}
	ctx.Logger.Infof("GetStaffPubKey: %x\n", giverPK)
	badge.Giver = giverPK

	badgeBytes, err := json.Marshal(badge)
	if err != nil {
		return nil, index, err
	}

	txID, err := ctx.broadcastTX("GiveBadge", badgeBytes, giverPK, sk, index, peers, key)
	if err != nil {
		return txID, index, err
	}

	index = (index + 1) % uint64(len(peers))

	return txID, index, nil
}

// ApproveActivity hashing activity, broadcast it and update to database
func (ctx *Context) ApproveActivity(activity *model.AttendedActivity, sk string, index uint64, peers []string, key []byte) (uint64, error) {
	approverPK, err := ctx.Database.GetStaffPublicKey(ctx.User.UserID)
	if err != nil {
		return index, err
	}

	activity.Approver = approverPK
	ctx.Logger.Infof("Publickey: %x\n", approverPK)

	approveBytes, err := json.Marshal(activity)
	if err != nil {
		return index, err
	}

	txID, err := ctx.broadcastTX("ApproveActivity", approveBytes, approverPK, sk, index, peers, key)
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

func (ctx *Context) broadcastTX(method string, params, pubKey []byte, privKey string, index uint64, peers []string, key []byte) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	// decrypting aes sk
	decSK, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}

	iv, err := hex.DecodeString("00112233445566778899aabbccddeeff")
	if err != nil {
		return nil, err
	}

	ctx.Logger.Infof("iv: %x (%d)\n", iv, len(iv))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decSK, decSK)

	ctx.Logger.Infof("decrypted aes sk: %x (%d)", string(decSK), len(decSK))
	ctx.Logger.Infof("sk [bytes]: %v\n", decSK)

	// Sign
	signature := ed25519.Sign(decSK, params)

	ctx.Logger.Infof("Sign: %s\nWith sk: 0x%x\nSignature: 0x%x\nTest: %v\n", params, decSK, signature, ed25519.Verify(pubKey, params, signature))

	payload := protoTm.Payload{
		Method: method,
		Params: params,
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
