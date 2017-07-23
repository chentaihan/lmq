package lmq

import (
	"fmt"
	"io/ioutil"
	"strings"

	"lmq/util"
	"lmq/container"
	"lmq/util/logger"
	"encoding/json"
)

const (
	MaxQueueCount = 100
	PlatformPath = "./data/platform/"
)

var platformManager PlatformManager

type PlatformManager struct{
	PlatformList container.Array
}

func initPlatform(){
	platformManager = PlatformManager{PlatformList: container.NewArray()}
	util.MkDir(PlatformPath)
	loadPlatform()
	jsonStr,_ := json.Marshal(platformManager.PlatformList)
	logger.Logger.Trace("InitPlatformManager success content="+string(jsonStr))
}

func loadPlatform(){
	dirList, _ := ioutil.ReadDir(PlatformPath)
	for _, file := range dirList{
		if !file.IsDir() {
			fileName := file.Name()
			if strings.HasSuffix(fileName, ".json"){
				platformManager.PlatformList.Append(fileName[0:len(fileName)-5])
			}
		}
	}
}

func AddPlatform(name string) bool{
	if !ExistPlatform(name){
		platformManager.PlatformList.Append(name)
		isOK := true
		for i := 0; i < 3; i++{
			isOK := util.CreateFile(GetPlatformPath(name))
			if isOK {
				break
			}
		}
		if !isOK {
			logger.Logger.Trace("AddPlatform failed platformName="+ name)
		}
		return isOK
	}
	return false
}

func GetPlatformPath(platform string) string{
	return fmt.Sprintf("%s%s.json", PlatformPath, platform)
}

func ExistPlatform(name string) bool{
	index := platformManager.PlatformList.Find(name);
	return index >= 0
}

func DeletePlatform(name string) bool{
	if item := platformManager.PlatformList.DeleteItem(name); item != nil{
		filePath := GetPlatformPath(name)
		return util.Remove(filePath)
	}
	return false
}

func SavePlatform(name string, content []byte) bool{
	filePath := GetPlatformPath(name)
	return util.WriteBytesCover(filePath, content)
}