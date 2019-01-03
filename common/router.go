package common

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"github.com/gorilla/mux"
)

var rt *mux.Router
var o sync.Once

func Router() *mux.Router {
	o.Do(func() {
		rt = mux.NewRouter()

		rt.HandleFunc("/oss-upload", upload2OssFunc)
	})
	return rt
}

func upload2OssFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		errMsg := "upload request method error, please use POST."
		SimpleResponseForErr(errMsg, w)
	} else {

	}
}

func SimpleResponseForErr(errMsg string, w io.Writer) {
	response := map[string]interface{}{
		"status": 401.1,
		"result": errMsg,
		"error":  "client error",
	}
	jso := jsonapiobj.JsResult{}
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
}
