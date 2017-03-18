package web

import (
	"goji.io"
	"goji.io/pat"
	"net/http"
)

type Handler func(request *http.Request) Responder
type gojiHandler func(w http.ResponseWriter, r *http.Request)

type Controller struct {
	mux     *goji.Mux
	headers map[string]string
}

func NewController() *Controller {
	con := new(Controller)
	con.mux = goji.SubMux()
	con.headers = make(map[string]string)
	return con
}

func (con *Controller) Mount(path string, root *goji.Mux) {
	root.Handle(pat.New(path), con.mux)
}

func (con *Controller) Route(routes *map[goji.Pattern]Handler) {
	for route, handler := range *routes {
		con.mux.HandleFunc(
			route, con.gojiHandlerFor(handler),
		)
	}
}

func (con *Controller) Header(key string, value string) {
	con.headers[key] = value
}

func (con *Controller) gojiHandlerFor(handler Handler) gojiHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responder := handler(r)
		addHeaders(w, con.headers)
		addHeaders(w, responder.Headers())
		responder.RespondOn(w)
	}
}

func addHeaders(w http.ResponseWriter, headers map[string]string) {
	for header, value := range headers {
		w.Header().Add(header, value)
	}
}
