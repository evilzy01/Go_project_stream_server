package session

// 1.拉session，当服务起来的时候，从db里边调 2. 新用户注册的时候，分配session id 3.校验的时候，看session过期了没，看用户是否合法登陆

import (
	"GO语言实战流媒体视频网站/api/dbops"
	"GO语言实战流媒体视频网站/api/defs"
	"GO语言实战流媒体视频网站/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map //  并发读性能非常好，写可能会需要加全局锁

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DelSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()     //毫秒
	ttl := ct + 30*60*1000 //过期时间定义为30min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InserSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true // 过期
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true //不OK——load出问题

}
