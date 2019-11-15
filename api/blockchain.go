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

	var giveBadgeRequest *model.GiveBadgeRequest
	if err := json.Unmarshal(body, &giveBadgeRequest); err != nil {
		ctx.Logger.Errorln("error while unmarshaling")
		return err
	}
	ctx.Logger.Infof("%+v", giveBadgeRequest)

	messageBytes, err := json.Marshal(giveBadgeRequest.Badges)
	if err != nil {
		ctx.Logger.Errorln("error when marshaling badges")
		return err
	}

	b64PubKey, err := ctx.Database.GetStaffPublicKey(ctx.User.UserID)
	if err != nil {
		ctx.Logger.Errorln("error when requesting publickey")
		return err
	}

	// Verify step
	isVerified, err := ctx.VerifySignature(messageBytes, giveBadgeRequest.Signature, string(b64PubKey))
	if err != nil {
		ctx.Logger.Errorln("unauthenticated in verifying")
		return err
	}

	if !isVerified {
		return errors.New("unauthenticated")
	}

	txs := make([]string, 0)
	for _, badge := range giveBadgeRequest.Badges {
		txID, currentIndex, err := ctx.GiveBadge(&badge, a.App.CurrentPeerIndex, a.Config.Peers, a.App.SK)
		a.App.CurrentPeerIndex = currentIndex
		if err != nil {
			return err
		}

		txs = append(txs, hex.EncodeToString(txID))
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
	if err := json.Unmarshal(body, &activityRequest); err != nil {
		return err
	}

	messageBytes, err := json.Marshal(activityRequest.Activities)
	if err != nil {
		return err
	}

	b64PubKey, err := ctx.Database.GetStaffPublicKey(ctx.User.UserID)
	if err != nil {
		return err
	}

	// Verify step
	isVerified, err := ctx.VerifySignature(messageBytes, activityRequest.Signature, string(b64PubKey))
	if err != nil {
		return err
	}

	if !isVerified {
		return errors.New("unauthenticated")
	}

	txs := make([]string, 0)
	for _, activity := range activityRequest.Activities {
		txID, currentIndex, err := ctx.ApproveActivity(&activity, a.App.CurrentPeerIndex, a.Config.Peers, a.App.SK)
		a.App.CurrentPeerIndex = currentIndex
		if err != nil {
			return err
		}

		txs = append(txs, hex.EncodeToString(txID))
	}

	txsBytes, err := json.Marshal(txs)
	if err != nil {
		return err
	}

	w.Write(txsBytes)

	return nil
}

// VerifyTX verifies whethear a corresponding
func (a *API) VerifyTX(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	ctx.Logger.Infof("json data: %s\n", body)

	isExists, evidence, currentIndex, err := ctx.VerifyTX(body, a.App.CurrentPeerIndex, a.Config.Peers)
	a.App.CurrentPeerIndex = currentIndex
	if err != nil {
		return err
	}

	mapper := make(map[string]interface{})
	mapper["result"] = isExists
	mapper["evidence"] = evidence

	mapperBytes, err := json.Marshal(mapper)
	if err != nil {
		return err
	}

	w.Write(mapperBytes)

	return nil
}
