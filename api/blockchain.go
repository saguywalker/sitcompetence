package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
		currentIndex, err := ctx.GiveBadge(badge, w, a.App.CurrentPeerIndex, a.Config.Peers)
		if err != nil {
			return err
		}

		a.App.CurrentPeerIndex = currentIndex
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
		currentIndex, err := ctx.ApproveActivity(activity, w, a.App.CurrentPeerIndex, a.Config.Peers)
		if err != nil {
			return err
		}

		a.App.CurrentPeerIndex = currentIndex
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

	rawData, err := json.Marshal(bodyMap["data"])
	if err != nil {
		return err
	}
	ctx.Logger.Infof("json data\n%s\n", rawData)

	isExists, currentIndex, hashData, err := ctx.VerifyTX(rawData, a.App.CurrentPeerIndex, a.Config.Peers)
	if err != nil {
		return err
	}

	a.App.CurrentPeerIndex = currentIndex

	returnResult := []string{fmt.Sprintf("Data hash(%x) was", hashData)}
	if !isExists {
		returnResult = append(returnResult, "not")
	}

	returnResult = append(returnResult, "recorded on the Blockchain.")

	result := strings.Join(returnResult, " ")

	if _, err := w.Write([]byte(result)); err != nil {
		return err
	}

	return nil
}
