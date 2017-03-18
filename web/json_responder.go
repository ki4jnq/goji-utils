package web

import (
	"encoding/json"
	"net/http"
)

type JsonResponder struct {
	data    interface{} // Arbitrary struct type
	headers map[string]string
	status  int
}

func NewJsonResponder(data interface{}) JsonResponder {
	jr := JsonResponder{data: data}
	jr.headers = make(map[string]string)
	jr.status = 200
	return jr
}

func (jr JsonResponder) RespondOn(w http.ResponseWriter) {
	jsonB, _ := json.Marshal(jr.data)
	w.Write([]byte(jsonB))
}

func (jr JsonResponder) Header(key string, value string) {
	jr.headers[key] = value
}

func (jr JsonResponder) Headers() map[string]string {
	return jr.headers
}

func (jr JsonResponder) SetStatus(code int) {
	jr.status = code
}

func (jr JsonResponder) Status() int {
	return jr.status
}
