package main

import (
	"GO语言实战流媒体视频网站/api/dbops"
	"GO语言实战流媒体视频网站/api/defs"
	"GO语言实战流媒体视频网站/api/session"
	"encoding/json"

	//	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreatUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDB)
		return
	}
	id := session.GenerateSessionId(ubody.Username)
	su := defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}

//func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	uname := p.ByName("user_name")
//	io.WriteString(w, uname)
//}
