package web

import (
	"net/http"
)

type StringResponder struct {
	response string
	headers  map[string]string
	status   int
}

func NewStringResponder(r string) Responder {
	return StringResponder{
		response: r,
	}
}

func (r StringResponder) RespondOn(w http.ResponseWriter) {
	w.Write([]byte(r.response))
}

func (r StringResponder) Header(key string, value string) {
	r.headers[key] = value
}

func (r StringResponder) Headers() map[string]string {
	return r.headers
}

func (r StringResponder) SetStatus(code int) {
	r.status = code
}

func (r StringResponder) Status() int {
	return r.status
}
