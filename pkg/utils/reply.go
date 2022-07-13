package utils

import (
	"encoding/json"
	"net/http"
)

func ReplyToReq(rw http.ResponseWriter, status int, data interface{}) {

	jsbody, _ := json.Marshal(data)
	rw.WriteHeader(status)
	rw.Write(jsbody)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
}
