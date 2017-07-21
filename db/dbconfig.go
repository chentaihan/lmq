package db

import(
	"encoding/json"
	"fmt"

	"lmq/util"
)
const(
	DBConf = "./data/db.conf"
	IndexFileName = "./data/index_%d.fdx"
	MessageFileName = "./data/message_%d.rdb"
	MaxRecordCount = 10000
)

type DBConfig struct{
	IdIndex int64
	indexFileName map[int64]string
	messageFileName map[int64]string
}

var dbConfig DBConfig

func InitConfig(){
	if conf ,_ := util.ReadFile(DBConf); len(conf) > 0 {
		json.Unmarshal(conf, &dbConfig)
	}
	dbConfig.indexFileName = make(map[int64]string)
	dbConfig.messageFileName = make(map[int64]string)
	index := getFileIndex(dbConfig.IdIndex)
	for i := index + 10; i >= index; i-- {
		dbConfig.indexFileName[i] = fmt.Sprintf(IndexFileName, i)
		dbConfig.messageFileName[i] = fmt.Sprintf(MessageFileName, i)
	}
}

func (dbConfig *DBConfig) Save() bool{
	if bytes, err := json.Marshal(dbConfig); err == nil{
		return util.WriteBytesCover(DBConf, bytes)
	}
	return false
}

func getFileIndex(index int64) int64{
	return index / MaxRecordCount
}

func (dbConfig DBConfig) GetIndexFileName(index int64) string{
	index = getFileIndex(index)
	if fileName, err := dbConfig.indexFileName[index]; err {
		return fileName;
	}
	fileName := fmt.Sprintf(IndexFileName, index)
	dbConfig.indexFileName[index] = fileName
	return fileName
}

func (dbConfig DBConfig) GetMessageFileName(index int64) string{
	index = getFileIndex(index)
	if fileName, err := dbConfig.messageFileName[index]; err {
		return fileName;
	}
	fileName := fmt.Sprintf(IndexFileName, index)
	dbConfig.messageFileName[index] = fileName
	return fileName
}

