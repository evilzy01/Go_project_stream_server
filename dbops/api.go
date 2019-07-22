package dbops

//在导入 mysql 驱动时, 这里使用了匿名导入的方式(在包路径前添加 _)
//当导入了一个数据库驱动后, 此驱动会自行初始化并注册自己到Golang的database/sql上下文中,
//因此我们就可以通过 database/sql 包提供的方法访问数据库了.
import (
	"GO语言实战流媒体视频网站/api/defs"
	"GO语言实战流媒体视频网站/api/utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//增删改查
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?,?)") //预编译
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error:%s", err)
		return err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create UUID
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // M D Y, HH:MM:SS要记住这个时间点——go语言的原点时间
	stmsIns, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmsIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DispalyCtime: ctime}
	defer stmsIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmsOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	var aid int
	var name string
	var dit string
	err = stmsOut.QueryRow(vid).Scan(&aid, &name, &dit)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmsOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DispalyCtime: dit}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	//create UUID
	stmsDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmsDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmsDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	StmsIns, err := dbConn.Prepare("INSERT INTO comments(id,video_id, author_id,content) VALUES(?,?,?,?)")
	if err != nil {
		return nil
	}
	_, err = StmsIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer StmsIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comments, error) {
	stmsOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content 
									FROM comments INNER JOIN users ON comments.author_id = users.id
									WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?) `)
	var res []*defs.Comments
	fmt.Println(vid)
	//	stmsOut.Query(vid, from, to)
	rows, err := stmsOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comments{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmsOut.Close()
	return res, nil
}
