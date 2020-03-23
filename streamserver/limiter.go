package main

/**
流量控制
*/

import (
  "log"
)

type ConnLimiter struct {
  concurrentConn int
  bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
  return &ConnLimiter{
    concurrentConn: cc,
    bucket:         make(chan int, cc), // 缓冲长度
  }
}

func (cl *ConnLimiter) GetConn() bool {
  log.Print(len(cl.bucket))
  if len(cl.bucket) >= cl.concurrentConn {
    log.Printf("Reached the rate limitation.")
    return false
  }
  cl.bucket <- 1
  return true
}

func (cl *ConnLimiter) ReleaseConn() {
  c :=<- cl.bucket
  log.Printf("New connection coming: %d", c)
}
