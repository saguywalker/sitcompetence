package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/cbergoon/merkletree"
	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GiveBadge takes a giving badge request and response with transactionID
func (a *API) GiveBadge(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
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

	// Struct of objects
	var listOfBadges []*model.GiveBadge
	if err := json.Unmarshal(decoded, &listOfBadges); err != nil {
		return err
	}

	// Hash of each object will be used to make a merkle tree
	listOfHashes := make([]merkletree.Content, len(listOfBadges))
	for i, x := range listOfBadges {
		hash, err := x.CalculateHash()
		if err != nil {
			return err
		}

		listOfHashes[i] = model.MyHash(fmt.Sprintf("%x", hash))
	}

	tree, err := merkletree.NewTree(listOfHashes)
	if err != nil {
		return err
	}

	merkleRootHash := fmt.Sprintf("%x", tree.MerkleRoot())
	ctx.Logger.Infof("Merkle Root Hash: %s", merkleRootHash)
	/*
		hashes := make([]string, uint(2*math.Ceil(float64(len(listOfHashes))/2.0)))
		for i, leaf := range tree.Leafs {
			hashes[i] = fmt.Sprintf("%x", leaf.Hash)
		}
	*/
	transactionID, err := a.broadcastTX(ctx, w, merkleRootHash)
	if err != nil {
		return err
	}

	if err := ctx.UpdateMerkleTransaction(transactionID, merkleRootHash, listOfHashes); err != nil {
		return err
	}

	if err := ctx.UpdateCollectedCompetence(listOfBadges, transactionID); err != nil {
		return err
	}

	return nil
}

// ApproveActivity takes an approving activity request and response with transactionID
func (a *API) ApproveActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	vals := r.URL.Query()
	params, ok := vals["data"]
	if !ok {
		return fmt.Errorf("missing data in parameter")
	}

	decoded, err := base64.StdEncoding.DecodeString(params[0])
	if err != nil {
		return err
	}

	var listOfActivities []model.ApproveActivity
	if err := json.Unmarshal(decoded, &listOfActivities); err != nil {
		return err
	}

	listOfHashes := make([]merkletree.Content, len(listOfActivities))
	for i, x := range listOfActivities {
		hash, err := x.CalculateHash()
		if err != nil {
			return err
		}

		listOfHashes[i] = model.MyHash(fmt.Sprintf("%x", hash))
	}

	// Free listOfActivities
	listOfActivities = nil

	tree, err := merkletree.NewTree(listOfHashes)
	if err != nil {
		return err
	}

	merkleRootHash := fmt.Sprintf("%x", tree.MerkleRoot())

	hashes := make([]string, uint(2*math.Ceil(float64(len(listOfHashes))/2.0)))
	for i, leaf := range tree.Leafs {
		hashes[i] = fmt.Sprintf("%x", leaf.Hash)
	}

	transactionID, err := a.broadcastTX(ctx, w, merkleRootHash)
	if err != nil {
		return nil
	}

	err = ctx.UpdateMerkleTransaction(transactionID, merkleRootHash, listOfHashes)

	return err
}

// VerifyTX verifies whethear a corresponding
func (a *API) VerifyTX(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	vals := r.URL.Query()
	txid, ok := vals["tx"]
	if !ok {
		return fmt.Errorf("missing transaction id parameter (tx)")
	}
	ctx.Logger.Infof("TxID: %s\n", string(txid[0]))

	data, ok := vals["data"]
	if !ok {
		return fmt.Errorf("missing json data parameter (data)")
	}

	decoded, err := base64.StdEncoding.DecodeString(data[0])
	if err != nil {
		return err
	}
	ctx.Logger.Infof("decoded value\n%s\n", string(decoded))

	return nil
}

// BroadcastTX broadcast transaction to blockchain node
// Note: Need to check wheater calling's node is reachable or not
func (a *API) broadcastTX(ctx *app.Context, w http.ResponseWriter, hash string) (string, error) {
	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=\"%x\"", a.Config.Peers[a.CurrentPeerIndex], hash)
	ctx.Logger.Infoln(url)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var fullData map[string]interface{}
	err = json.Unmarshal(data, &fullData)
	if err != nil {
		return "", err
	}

	resultData, ok := fullData["result"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("error when asserting type from data[\"result\"]")
	}

	txIDInterface := resultData["hash"]
	transactionID := strings.ToLower(fmt.Sprintf("%v", txIDInterface))
	ctx.Logger.Infof("TransactionID: %s\n", transactionID)

	//w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)

	// Move to the next peer in round-robin fashion
	a.CurrentPeerIndex = (a.CurrentPeerIndex + 1) % len(a.Config.Peers)

	return transactionID, err
}
