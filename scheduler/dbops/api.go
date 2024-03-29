package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddVideoDelRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

// why we don't write this function in the internal.go
