package dbops

import (
	"GO语言实战流媒体视频网站/api/defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

// 1.写session 2.拿session信息 3.list所有session 4. 删除session

func InserSession(sid string, ttl int64, uname string) error {
	stmsIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	ttlstr := strconv.FormatInt(ttl, 10)
	_, err = stmsIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmsIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	stmsOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmsOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	ss := &defs.SimpleSession{}
	res, err := strconv.ParseInt(ttl, 10, 64)
	if err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmsOut.Close()
	return ss, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmsOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmsOut.Query()
	if err != nil {
		log.Panicf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var uname string
		err = rows.Scan(&id, &ttlstr, &uname)
		if err != nil {
			log.Panicf("retrieve sessions error: %s", err)
		}
		ttl, err1 := strconv.ParseInt(ttlstr, 10, 64)
		if err1 == nil {
			ss := &defs.SimpleSession{Username: uname, TTL: ttl}
			m.Store(id, ss)
			log.Panicf("session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

func DelSession(id string) error {
	stmsDel, err := dbConn.Prepare("DELETE FROM session WHERE session = ?")
	if err != nil {
		log.Panicf("%s", err)
		return err
	}
	if _, err := stmsDel.Exec(id); err != nil {
		return err
	}
	return nil
}
