package web

import "net/http"

type Responder interface {
	RespondOn(w http.ResponseWriter)

	Header(name string, value string)
	Headers() map[string]string

	SetStatus(code int)
	Status() int
}
