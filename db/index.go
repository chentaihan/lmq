package db

import (
	"time"
	"sync"
	"encoding/json"
	"fmt"

	"lmq/util/logger"
	"lmq/util"
	"lmq/container"
)

const (
	IndexLength = int64(76)

)

type Index struct {
	Index  int64
	Offset int64
	Length int64
}

func NewIndex() *Index {
	index := new(Index)
	index.Index = increment()
	return index
}

type IncrementId struct {
	id   int64
	lock sync.RWMutex
}

var incrementId IncrementId
var indexMap *container.SyncMap

func InitIndex() {
	incrementId.id = dbConfig.IdIndex
	indexMap = container.NewSyncMap()
}

func getIndex() int64 {
	incrementId.lock.RLock()
	defer incrementId.lock.RUnlock()
	return incrementId.id
}

func increment() int64 {
	incrementId.lock.Lock()
	defer incrementId.lock.Unlock()
	incrementId.id++
	dbConfig.IdIndex = incrementId.id
	return incrementId.id
}

func SaveIndex(index *Index) bool{
	buf := make([]byte, IndexLength)
	strBuf, _ := json.Marshal(*index)
	for i := 0; i < len(strBuf); i++ {
		buf[i] = strBuf[i]
	}
	buf[IndexLength-2] ='\r'
	buf[IndexLength-1] ='\n'
	offset := (index.Index - 1) * IndexLength
	indexMap.Add(index.Index, index)
	isSuccess := util.WriteFileOffset(dbConfig.GetIndexFileName(index.Index), offset, buf)
	if isSuccess {
		isSuccess = dbConfig.Save()
	}else{
		logger.Logger.Errorf("SaveIndex failed %s", buf)
	}
	return isSuccess
}

func deleteIndex(index int64, nameIndex uint){
	logger.Logger.Tracef("deleteIndex id=%d, nameIndex=%d", index, nameIndex)
	indexMap.Delete(index)
}

func GetIndex(id int64, nameIndex uint) *Index {
	logger.Logger.Tracef("GetIndexInfo id=%d, nameIndex=%d", id, nameIndex)
	if idxInfo := indexMap.Get(id); idxInfo != nil{
		ret, _ := idxInfo.(*Index)
		return ret
	}

	offset := (id - 1) * IndexLength
	fileName := dbConfig.GetIndexFileName(id)
	bytes, _ := util.ReadFileOffset(fileName, offset, IndexLength)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 0 {
			bytes = bytes[0:i]
			break
		}
	}
	var indexInfo Index
	if err := json.Unmarshal(bytes, &indexInfo); err != nil{
		logger.Logger.Errorf("get from database failed %d", id)
		return nil
	}
	indexMap.Add(id, indexInfo)
	return &indexInfo;
}

func IndexTest() {
	startTime := time.Now().Unix()

	for i := 0; i < 10;i++  {
		go func(){
			var i int64 = 0
			for ; i < 100; i++ {
				id := increment()
				indexInfo := NewIndex()
				indexInfo.Index = id;
				indexInfo.Offset = 10*i
				indexInfo.Length = 10*i
				SaveIndex(indexInfo);
			}
		}()
	}

	GetIndex(5,0)
	deleteIndex(5,0)
	GetIndex(5,0)
	GetIndex(107,0)
    endTime := time.Now().Unix()
	fmt.Println("time: ",endTime-startTime)
}