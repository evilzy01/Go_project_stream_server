package main

import (
	"net/http"

	"GO语言实战流媒体视频网站/scheduler/dbops"

	"github.com/julienschmidt/httprouter"
)

func vidDelHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(w, "video is should not be empty", 400)
		return
	}

	err := dbops.AddVideoDelRecord(vid)
	if err != nil {
		sendResponse(w, "Internal server error", 500)
		return
	}

	sendResponse(w, "", 200)
	return
}
