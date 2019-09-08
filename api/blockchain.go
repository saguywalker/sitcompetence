package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/saguywalker/sitcompetence/app"
)

func (a *API) BroadcastTX(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	logrus.Infoln("In BroadcastTX function.")

	tx := mux.Vars(r)["tx"]

	if !strings.HasPrefix(tx, "val:") || len(bytes.Split([]byte(tx), []byte("="))) != 2 {
		return fmt.Errorf("invalid transaction format (got %v)", tx)
	}

	url := fmt.Sprintf("http://%s:26657/broadcast_tx_commit?tx=\"%s\"", a.Config.Peers[a.CurrentPeerIndex], tx)
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	_, err = w.Write(data)

	return err
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
