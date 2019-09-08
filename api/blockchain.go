package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/saguywalker/sitcompetence/app"
)

func (a *API) BroadcastTX(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	logrus.Infoln("In BroadcastTX function.")

	vals := r.URL.Query()
	tx, ok := vals["tx"]
	if !ok {
		return fmt.Errorf("missing tx parameter")
	}

	if len(tx) != 1 {
		return fmt.Errorf("invalid tx format")
	}

	logrus.Infoln(tx[0])

	if !strings.HasPrefix(tx[0], "val:") && len(bytes.Split([]byte(tx[0]), []byte("="))) != 2 {
		return fmt.Errorf("invalid transaction format (got %v)", tx[0])
	}

	url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=%s", a.Config.Peers[a.CurrentPeerIndex], tx[0])
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
