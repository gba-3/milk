package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type JsonHandler func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)

func (jh JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := jh(w, r)
	if err != nil {
		responseJson(w, http.StatusBadRequest, res)
	}
	responseJson(w, status, res)

}

func responseJson(w http.ResponseWriter, status int, payload interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}
