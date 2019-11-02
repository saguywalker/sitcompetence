package api

import (
	"encoding/hex"
	"encoding/json"
	"errors"
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
		return errors.New("giveBadge must be called by admin only")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	ctx.Logger.Infof("%s\n", body)

	// Struct of objects
	var giveBadgeRequest *model.GiveBadgeRequest
	// var listOfBadges []*model.CollectedCompetence
	if err := json.Unmarshal(body, &giveBadgeRequest); err != nil {
		ctx.Logger.Errorln("error while unmarshaling")
		return err
	}

	ctx.Logger.Infof("%+v", giveBadgeRequest)

	txs := make([]string, 0)

	for _, badge := range giveBadgeRequest.Badges {
		txID, currentIndex, err := ctx.GiveBadge(&badge, giveBadgeRequest.PrivateKey, a.App.CurrentPeerIndex, a.Config.Peers, a.App.Config.SecretKey)
		if err != nil {
			return err
		}

		txs = append(txs, hex.EncodeToString(txID))
		a.App.CurrentPeerIndex = currentIndex
	}

	txsBytes, err := json.Marshal(txs)
	if err != nil {
		return err
	}

	w.Write(txsBytes)

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

	var activityRequest *model.ApproveActivityRequest
	// var listOfActivities []*model.AttendedActivity
	if err := json.Unmarshal(body, &activityRequest); err != nil {
		return err
	}

	for _, activity := range activityRequest.Activities {
		currentIndex, err := ctx.ApproveActivity(&activity, activityRequest.PrivateKey, a.App.CurrentPeerIndex, a.Config.Peers, a.App.Config.SecretKey)
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
	/*
		var bodyMap map[string]interface{}
		if err := json.Unmarshal(body, &bodyMap); err != nil {
			return nil
		}

		rawData, err := json.Marshal(bodyMap["data"])
		if err != nil {
			return err
		}
		ctx.Logger.Infof("json data: %s\n", rawData)
	*/
	ctx.Logger.Infof("json data: %s\n", body)

	isExists, currentIndex, err := ctx.VerifyTX(body, a.App.CurrentPeerIndex, a.Config.Peers)
	if err != nil {
		return err
	}

	a.App.CurrentPeerIndex = currentIndex

	w.Write([]byte(fmt.Sprintf("%v", isExists)))

	return nil
}
