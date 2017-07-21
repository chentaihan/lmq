package db

import (
	"encoding/json"

	"lmq/util"
	"lmq/util/logger"
)

type Message struct{
	ID int64
	Platform string
	Module string
	Tag string
	Url string
	Params string
}

func SaveMessage(msg *Message) int64{
	indexInfo := NewIndex()
	msg.ID = indexInfo.Index
	bytes, _ := json.Marshal(msg)
	fileName := dbConfig.GetMessageFileName(msg.ID);
	isSuccess := util.WriteBytes(fileName, bytes)
	if isSuccess {
		indexInfo.Length = int64(len(bytes))
		fileName = dbConfig.GetMessageFileName(indexInfo.Index)
		indexInfo.Offset = util.FileSize(fileName)
		isSuccess = SaveIndex(indexInfo)
		if isSuccess {
			return msg.ID
		}
	}else{
		logger.Logger.Errorf("SaveMessage failed %s", bytes)
	}
	return -1
}
