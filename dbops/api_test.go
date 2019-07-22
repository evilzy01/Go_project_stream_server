package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempid string

//////////////////////////////////////////////

func Testmain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

//func TestUserWorkflow(t *testing.T) {
//	t.Run("add", testAddUsers)
//	t.Run("get", testGetUsers)
//	t.Run("del", testDelUsers)
//	t.Run("Reget", testRegetUsers)
//}

////////////////////////////////////////
func testAddUsers(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of add users: %v", err)
	}
}

func testGetUsers(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of get users")
	}
}

func testDelUsers(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of del users: %v", err)
	}
}

func testRegetUsers(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of Reget users: %v", err)
	}
	if pwd != "" {
		t.Errorf("Del Failed")
	}
}

//func TestVideoWorkflow(t *testing.T) {
//	t.Run("PrepareUser", testAddUsers)
//	t.Run("add", testAddVideo)
//	t.Run("get", testGetVideo)
//	t.Run("del", testDelVideo)
//	t.Run("Reget", testRegetVideo)
//}

func testAddVideo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-first-video")
	if err != nil {
		t.Errorf("Error of add video: %v", err)
	}
	tempid = vi.Id
}

func testGetVideo(t *testing.T) {
	vi, err := GetVideoInfo(tempid)
	if vi.AuthorId != 1 || err != nil {
		t.Errorf("Error of get videoinfo: = %v", err)
	}
}

func testDelVideo(t *testing.T) {
	err := DeleteVideoInfo(tempid)
	if err != nil {
		t.Errorf("Error of del videos: %v", err)
	}
}

func testRegetVideo(t *testing.T) {
	vi, err := GetVideoInfo(tempid)
	if vi != nil || err != nil {
		t.Errorf("Error of reget videoinfo: = %v", err)
	}
}

func TestCommentWorkflow(t *testing.T) {
	t.Run("PrepareUser", testAddUsers)
	//	t.Run("add", testAddVideo)
	t.Run("addComments", testAddComments)
	t.Run("listComments", testListComments)
}

func testAddComments(t *testing.T) {
	err := AddNewComments(tempid, 1, "this is a good movie")
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	fmt.Println(to)
	cc, err := ListComments("12345", 1514764800, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range cc {
		fmt.Printf("comment : %d,%v,\n", i, ele)
	}
}
