package main

import (
	"GO语言实战流媒体视频网站/api/defs"
	"GO语言实战流媒体视频网站/api/session"

	//	"fmt"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

//  Session 校验
func validUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		//		fmt.Println("auth")
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

// User校验，不知道这个是干什么的????
func validUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser) /// 不需要输入参数么？？？？
		return false
	}
	return true
}
