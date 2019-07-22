package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/////////////// MiddleWare to add authentication ////////////
type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validUserSession(r)
	m.r.ServeHTTP(w, r) //  ServeHTTP is a built in function___makes the router implement the http.Handler interface.
}

func NewMiddelWareHandler(r *httprouter.Router) http.Handler { // http.Handler is a interface
	m := middleWareHandler{}
	m.r = r
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreatUser)
	//	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddelWareHandler(r)
	http.ListenAndServe(":8006", mh)
}
