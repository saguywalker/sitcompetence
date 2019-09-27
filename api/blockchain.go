package api

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GiveBadge takes a giving badge request and response with transactionID
func (a *API) GiveBadge(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// Struct of objects
	var listOfBadges []*model.CollectedCompetence
	if err := json.Unmarshal(body, &listOfBadges); err != nil {
		return err
	}

	for _, badge := range listOfBadges {
		err := ctx.GiveBadge(badge, w, a.Config.Peers)
		if err != nil {
			return err
		}
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

	var listOfActivities []*model.AttendedActivity
	if err := json.Unmarshal(body, &listOfActivities); err != nil {
		return err
	}

	for _, activity := range listOfActivities {
		err := ctx.ApproveActivity(activity, w, a.Config.Peers)
		if err != nil {
			return err
		}
	}

	return nil
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

	decodedTxid, err := base64.StdEncoding.DecodeString(bodyMap["transaction_id"].(string))
	if err != nil {
		return err
	}
	ctx.Logger.Infof("decoded transaction id: %x", decodedTxid)

	rawData, err := json.Marshal(bodyMap["data"])
	if err != nil {
		return err
	}
	ctx.Logger.Infof("json data\n%s\n", rawData)

	result, err := ctx.VerifyTX(rawData, decodedTxid)
	if err != nil {
		return err
	}

	var returnResult string
	if result {
		returnResult = fmt.Sprintf("Data was recorded in blockchain at TransactionID: %v :)", bodyMap["transaction_id"])
	} else {
		returnResult = fmt.Sprintf("Data was not recorded in the blockchain :(.")
	}

	if _, err := w.Write([]byte(returnResult)); err != nil {
		return err
	}

	return nil
}

// BroadcastTX broadcast transaction to blockchain node
// Note: Need to check wheater calling's node is reachable or not
func (a *API) broadcastTX(ctx *app.Context, w http.ResponseWriter, hash []byte) ([]byte, error) {
	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", a.Config.Peers[ctx.CurrentPeerIndex], hash)
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

	// w.WriteHeader(http.StatusOK)
	// if _, err := w.Write(data); err != nil {
	// panic(err)
	// }

	// Move to the next peer in round-robin fashion
	ctx.CurrentPeerIndex = (ctx.CurrentPeerIndex + uint64(1)) % uint64(len(ctx.Peers))

	return transactionID, err
}
