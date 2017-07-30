package event

import (
	"fmt"
	"strconv"
	"time"
	"sync"

	"lmq/container"
	"lmq/util/logger"
	"lmq/lmq"
)

const (
	EVENT_TYPE_ADD_MESSAGE = 0;
	EVENT_TYPE_DELETE_QUEUE = 2;
)

var msgCenter *EventCenter

type EventItem struct{
	Chan chan int
	Queue *container.CQueue
}

type EventCenter struct{
	rwLock sync.RWMutex
	MsgQueue map[string]*EventItem
}

func InitEventCenter(moduleNameList []string, esQueueList []*container.CQueue) {
	msgCenter = new(EventCenter)
	msgCenter.MsgQueue = make(map[string]*EventItem, len(moduleNameList))
	for i := len(moduleNameList)-1; i >= 0; i--{
		eventItem := &EventItem{Chan: make(chan int, 64), Queue: esQueueList[i]}
		msgCenter.MsgQueue[moduleNameList[i]] = eventItem
	}
	logger.Logger.Tracef("InitEventCenter success, moduleCount=%d", len(moduleNameList))
}

func AddQueue(queueName string, queue *container.CQueue) bool{
	logger.Logger.Tracef("AddQueue queueName=%s", queueName)
	msgCenter.rwLock.RLock()
	defer msgCenter.rwLock.Unlock()
	if _,ok := msgCenter.MsgQueue[queueName]; !ok {
		eventItem := &EventItem{Chan: make(chan int, 64), Queue: queue}
		msgCenter.MsgQueue[queueName] = eventItem
		startEventItem(queueName, eventItem)
		return true
	}
	return false
}

func DeleteQueue(queueName string) bool{
	logger.Logger.Tracef("DeleteQueue queueName=%s", queueName)
	msgCenter.rwLock.RLock()
	defer msgCenter.rwLock.Unlock()
	if eventItem,ok := msgCenter.MsgQueue[queueName]; !ok {
		eventItem.Chan <- EVENT_TYPE_DELETE_QUEUE
		return true
	}
	return false
}

func StartEventCenter(){
	for queueName, eventItem := range msgCenter.MsgQueue{
		startEventItem(queueName, eventItem)
	}
	logger.Logger.Tracef("StartEventCenter success")
}

func startEventItem(queueName string, eventItem *EventItem){
	go func() {
		for{
			var eventType int
			select {
			case eventType = <- eventItem.Chan:
				logger.Logger.Tracef("startEventItem eventType=%d, queueSize=%d", eventType, eventItem.Queue.Size())
				if eventType == EVENT_TYPE_ADD_MESSAGE {
					msgQueue := eventItem.Queue.Copy()
					if msgQueue != nil{
						for i := 0; i < len(msgQueue); i++ {
							msg,_ := msgQueue[i].(*lmq.Message)
							msg.Execute()
						}
					}
				}
			}
			if eventType == EVENT_TYPE_DELETE_QUEUE {
				msgCenter.rwLock.RLock()
				close(eventItem.Chan)
				delete(msgCenter.MsgQueue, queueName)
				logger.Logger.Tracef("delete queueName=%s", queueName)
				msgCenter.rwLock.Unlock()
				break
			}
		}
	}()
}

func SendSignal(queueName string, signal int) bool{
	logger.Logger.Tracef("SendSignal queueName=%s, signal=%d", queueName,signal)
	if eventItem,ok := msgCenter.MsgQueue[queueName]; ok{
		eventItem.Chan <- signal
		return true
	}else{
		logger.Logger.Errorf("SendSignal queueName=%s, signal=%d not exist", queueName,signal)
		return false
	}
}

func TestEvent(){
	var ch1 chan int
	ch1 = make (chan int, 4)
	ch2 := make (chan int, 0)
	go func() {
		for{
			select {
			case val := <- ch1:
				fmt.Println("ch1 pop one element" + strconv.Itoa(val))
				time.Sleep(time.Second * 5)

			case <-ch2:
				fmt.Println("ch2 pop one element")
			}
		}
		close(ch1)
		close(ch2)
	}()
	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
}
