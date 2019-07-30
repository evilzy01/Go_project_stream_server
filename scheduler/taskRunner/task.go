package taskRunner

// content in the task.go is mostly specific, can't reuse.
// Here, we only focus on the task of our project_ delete delay
// 1. Read information that we want to delete in the database and put data in the data channel_ Dispatcher
// 2. delete the real source_ Executor

import (
	"GO语言实战流媒体视频网站/scheduler/dbops"
	"errors"
	"log"
	"os"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(vid + ".mp4")
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error : %v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChnn) error {
	res, err := dbops.ReadVideoDelRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChnn) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) { // why we dont use vid directly. Since func is a clouser，在一个goroutine中，它只能拿到瞬时的参数状态，不会保存，如果想用vid，需要当做参数传进来
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDelRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)

		default:
			break forloop
		}
	}
	errMap.Range(func(k, v interface{}) bool { // why we need to scan all the errors again?????
		err := v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
