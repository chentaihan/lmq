package event

import (
	"fmt"
	"strconv"
	"time"
	"lmq/container"
)

const (
	EVENT_TYPE_ADD = 0;
	EVENT_TYPE_DELETE = 2;
)

var msgCenter *EventCenter

type EventItem struct{
	Chan chan int
	Queue *container.EsQueue
}

type EventCenter struct{
	MsgQueue map[string]*EventItem
}

func InitEventCenter(queueCount int) {
	msgCenter = new(EventCenter)
	msgCenter.MsgQueue = make(map[string]*EventItem, queueCount)
	return msgCenter
}

func AddQueue(queueName string, queue *container.EsQueue) bool{
	if _,ok := msgCenter.MsgQueue[queueName]; !ok {
		eventItem := &EventItem{Chan: make(chan int, 64), Queue: queue}
		msgCenter.MsgQueue[queueName] = eventItem
		startEventItem(queueName, eventItem)
		return true
	}
	return false
}

func DeleteQueue(queueName string) bool{
	if eventItem,ok := msgCenter.MsgQueue[queueName]; !ok {
		eventItem.Chan <- EVENT_TYPE_DELETE
		return true
	}
	return false
}

func StartEventCenter(){
	for queueName, eventItem := range msgCenter.MsgQueue{
		startEventItem(queueName, eventItem)
	}
}

func startEventItem(queueName string, eventItem *EventItem){
	go func() {
		for{
			var eventType int
			select {
			case eventType <- eventItem.Chan:
				if eventType == EVENT_TYPE_ADD {
					//
				}
			}
			if eventType == EVENT_TYPE_DELETE {
				close(eventItem.Chan)
				delete(msgCenter.MsgQueue, queueName)
				break
			}
		}
	}()
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
