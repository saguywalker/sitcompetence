package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GiveBadge takes a giving badge request and response with transactionID
func (a *API) GiveBadge(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.Logger.Infoln(ctx.User.Group)
	if ctx.User.Group != "inst_group" {
		return fmt.Errorf("giveBadge must be called by admin only")
	}

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
		currentIndex, err := ctx.GiveBadge(badge, a.App.CurrentPeerIndex, a.Config.Peers)
		if err != nil {
			return err
		}

		a.App.CurrentPeerIndex = currentIndex
	}

	return nil
}

// ApproveActivity takes an approving activity request and response with transactionID
func (a *API) ApproveActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	if ctx.User.Group != "inst_group" {
		return fmt.Errorf("giveBadge must be called by admin only")
	}

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
		currentIndex, err := ctx.ApproveActivity(activity, a.App.CurrentPeerIndex, a.Config.Peers)
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

	isExists, currentIndex, err := ctx.VerifyTX(rawData, a.App.CurrentPeerIndex, a.Config.Peers)
	if err != nil {
		return err
	}

	a.App.CurrentPeerIndex = currentIndex

	w.Write([]byte(fmt.Sprintf("%v", isExists)))

	return nil
}
