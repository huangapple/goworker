package goworker

import (
	"fmt"
	"runtime"
	"time"
)

type jobQueue_ struct {
	queue     chan func()
	workQueue chan int //当前工作者
}

//实例化一个jobQueue
func newJobQueue(maxQueue int, maxWorker int) *jobQueue_ {

	jq := new(jobQueue_)

	jq.queue = make(chan func(), maxQueue)
	jq.workQueue = make(chan int, maxWorker)

	go func() {
		for {
			select {
			case jobInfo := <-jq.queue:
				jq.workQueue <- 1
				jq.DoJob(jobInfo)
			}
		}
	}()

	return jq
}

func (jq *jobQueue_) PushJob(job func()) {

	jq.queue <- job
}

func (jq *jobQueue_) GetJob() <-chan func() {
	return jq.queue
}

func (jq *jobQueue_) DoJob(job func()) {

	go func() {
		st := time.Now()
		defer func() {
			<-jq.workQueue
			if onInfoLog != nil {
				onInfoLog("任务执行时间：" + time.Now().Sub(st).String())
			}
			if r := recover(); r != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				if onErrorLog != nil {
					onErrorLog(fmt.Sprintf("worker panic : %s\n%s", r, buf))
				}
			}
		}()
		//执行job
		job()
	}()

}
