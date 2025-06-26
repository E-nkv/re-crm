package app

import (
	"encoding/json"
	"net/http"
	"re-crm/utils"
)

type httpWriter struct{}

func (httpw *httpWriter) Json(w http.ResponseWriter, code int, data utils.Object) {
	bs, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "error marshaling data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bs)
}

func (httpw *httpWriter) Error(w http.ResponseWriter, code int, msg string) {
	http.Error(w, msg, code)

}

type websocketWriter struct{}

func (ww *websocketWriter) Json(data utils.Object) {

}
