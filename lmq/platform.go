package lmq

import (
	"fmt"

	"lmq/util"
	"encoding/json"
	"lmq/container"
	"lmq/util/logger"
)

const (
	MaxQueueCount = 100
)

var PlatformList map[string]*Platform

type Platform struct{
	Name string
	ModuleList container.Array
	UsedQueueCount int
	MaxQueueCount int
}

type PlatformItem struct{
	Platform string
	Module string
}

func NewPlatform(name string) *Platform{
	return &Platform{Name : name}
}

func LoadPlatform() bool{
	if PlatformList == nil {
		PlatformList = make(map[string]*Platform)
	}
	if bytes := util.LoadJson(util.PlatformConf); bytes != nil{
		var platforms []PlatformItem
		json.Unmarshal(bytes, &platforms)
		for _, pf := range platforms{
			AddPlatform(pf)
		}
		return true
	}
	return false
}

func AddPlatform(pf PlatformItem) bool{
	platform, ok := PlatformList[pf.Platform];
	if !ok {
		PlatformList[pf.Platform] = NewPlatform(pf.Platform)
		platform = PlatformList[pf.Platform]
		platform.MaxQueueCount = MaxQueueCount
	}
	if index := platform.ModuleList.Find(pf.Module); index == -1{
		platform.ModuleList.Append(pf.Module)
		return true
	}else{
		logger.Logger.Errorf("%s %s is exist", pf.Platform, pf.Module)
		return  false
	}
}

func OutPutPlatformList(){
	str, err := json.Marshal(PlatformList);
	if err == nil{
		fmt.Println(string(str))
	}
}
