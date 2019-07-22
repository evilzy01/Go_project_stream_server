package main

import (
	"log"
)

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

// 构造函数

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	return true
}

func (cl *ConnLimiter) RealseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d", c)
}
