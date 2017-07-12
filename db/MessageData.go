package db

import (
	"encoding/json"
	"lmq/util"
	"lmq/container"
)

type Message struct{
	ID int64
	Platform string
	Module string
	Tag string
	Url string
	Params string
}

var fileSizeMap container.Map
var writeMessageChan chan *Message

func InitMessage(){
	writeMessageChan = make(chan *Message, 10)

	go func() {
		for {
			select {
			case message := <-writeMessageChan:
				SaveMessage(message)
			}
		}
		close(writeMessageChan)
	}()
}

func AddMessage(msg *Message){
	indexInfo := NewIndex()
	msg.ID = indexInfo.Index
	bytes, _ := json.Marshal(msg)
	indexInfo.Length = int64(len(bytes))
	if  offset := fileSizeMap.Get(getFileIndex(indexInfo.Index)); offset != nil{
		indexInfo.Offset,_ = offset.(int64)
	}else{
		fileName := dbConfig.GetMessageFileName(indexInfo.Index)
		indexInfo.Offset = util.FileSize(fileName)
	}
	fileSizeMap.Add(getFileIndex(indexInfo.Index), indexInfo.Length + indexInfo.Offset)
	writeMessageChan <- msg
	SaveIndex(indexInfo)
}

func SaveMessage(msg *Message){
	buf, _ := json.Marshal(*msg)
	fileName := dbConfig.GetMessageFileName(msg.ID);
	util.WriteBytes(buf, fileName)
}
