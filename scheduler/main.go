package main

import (
	"net/http"

	"GO语言实战流媒体视频网站/scheduler/taskRunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video_del_record/:vid-id", vidDelHandler)

	return router
}

func main() {
	go taskRunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
