package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

// 构造函数
func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, "Too many connections", http.StatusTooManyRequests)
	}
	m.r.ServeHTTP(w, r)
	defer m.l.RealseConn()

}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	r := RegisterHandler()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":8090", mh)
}
