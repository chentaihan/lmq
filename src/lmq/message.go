package lmq

import (
	"../util"
)

const(
	MESSAGE_STATUS_READY = 0;
	MESSAGE_STATUS_RUNNING = 1;
	MESSAGE_STATUS_FAILED = 2;
	MESSAGE_STATUS_SUCCESS = 3;
)

type Message struct{
	ID int64
	Platform string
	Module string
	Tag string
	Url string
	Params string
	Status int
}

func NewMessage() *Message{
	return &Message{Status: MESSAGE_STATUS_READY}
}

func (msg *Message) Execute()(string,error){
	return util.HttpPostJson(msg.Url, msg.Params)
}
