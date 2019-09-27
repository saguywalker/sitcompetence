package api

import (
	"encoding/base64"
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
