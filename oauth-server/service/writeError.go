package service

import (
	netHttp "net/http"

	"github.com/plgd-dev/kit/codec/json"
	"github.com/plgd-dev/kit/log"
)

func errToJsonRes(err error) map[string]string {
	return map[string]string{"err": err.Error()}
}

func writeError(w netHttp.ResponseWriter, err error, status int) {
	if err == nil {
		w.WriteHeader(netHttp.StatusNoContent)
		return
	}
	log.Errorf("%v", err)
	b, _ := json.Encode(errToJsonRes(err))
	netHttp.Error(w, string(b), status)
}
