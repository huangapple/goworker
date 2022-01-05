package goworker

import (
	"fmt"
	"sync"
)

var mqMap sync.Map
var onErrorLog func(string)
var onInfoLog func(string)

func InitLog(onError func(string), onInfo func(string)) {
	onErrorLog = onError
	onInfoLog = onInfo

}

//最大任务队列， 与最大工作线程
func Init(name string, maxQueue int, maxWorker int) {

	jobQueue := newJobQueue(maxQueue, maxWorker)
	mqMap.Store(name, jobQueue)
}

//丢任务进去
func Push(name string, job func()) {

	mqInter, ok := mqMap.Load(name)

	if !ok {
		if onErrorLog != nil {
			onErrorLog(fmt.Sprintf("name:%s的mq不存在", name))
		}
		job()
		return
	}

	mq, ok := mqInter.(*jobQueue_)
	if !ok {
		if onErrorLog != nil {
			onErrorLog(fmt.Sprintf("name:%s的mq不能转成 *jobQuene_", name))
		}
		job()
		return
	}

	mq.PushJob(job)
}
