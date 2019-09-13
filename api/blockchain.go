package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cbergoon/merkletree"
	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

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

	var listOfBadges []model.GiveBadge
	if err := json.Unmarshal(decoded, &listOfBadges); err != nil {
		return err
	}

	listOfContents := make([]merkletree.Content, len(listOfBadges))
	for i, x := range listOfBadges {
		listOfContents[i] = x
	}

	// Free listOfBadges
	listOfBadges = nil

	tree, err := merkletree.NewTree(listOfContents)
	if err != nil {
		return err
	}

	merkleRootHash := tree.MerkleRoot()

	hashes := make([][]byte, len(listOfContents))
	for i, leaf := range tree.Leafs {
		hashes[i] = leaf.Hash
	}

	transactionID, err := a.BroadcastTX(ctx, w, merkleRootHash)
	if err != nil {
		return nil
	}

	err = ctx.GiveBadge(transactionID, merkleRootHash, hashes)

	return err
}

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

	listOfContents := make([]merkletree.Content, len(listOfActivities))
	for i, x := range listOfActivities {
		listOfContents[i] = x
	}

	// Free listOfActivities
	listOfActivities = nil

	tree, err := merkletree.NewTree(listOfContents)
	if err != nil {
		return err
	}

	merkleRootHash := tree.MerkleRoot()

	hashes := make([][]byte, len(listOfContents))
	for i, leaf := range tree.Leafs {
		hashes[i] = leaf.Hash
	}

	transactionID, err := a.BroadcastTX(ctx, w, merkleRootHash)
	if err != nil {
		return nil
	}

	err = ctx.ApproveActivity(transactionID, merkleRootHash, hashes)

	return err
}

// BroadcastTX broadcast transaction to blockchain node
func (a *API) BroadcastTX(ctx *app.Context, w http.ResponseWriter, hash []byte) ([]byte, error) {
	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", a.Config.Peers[a.CurrentPeerIndex], hash)

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
		return nil, fmt.Errorf("error when asserting type from data[\"result\"]")
	}

	transactionID := resultData["hash"]
	transactionBytes, ok := transactionID.([]byte)
	if !ok {
		return nil, fmt.Errorf("error when asserting type from transactionID")
	}

	_, err = w.Write([]byte(fmt.Sprintf("%x", transactionBytes)))

	// Move to the next peer in round-robin fashion
	a.CurrentPeerIndex = (a.CurrentPeerIndex + 1) % len(a.Config.Peers)

	return transactionBytes, err
}

/*
func (a *API) ApproveActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.ApproveActivity

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.GiveBadge(body); err != nil {
		return err
	}

	data, err := json.Marshal(&CreateActivityResponse{ActivityID: input.ActivityID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}
*/
