package lmq

import (
	"fmt"
	"encoding/json"

	"lmq/container"
	"lmq/util"
	"lmq/util/logger"
)

type Module struct{
	Platform string
	Name     string
	Queue    *container.CQueue
}

var moduleManager ModuleManager

type ModuleManager struct{
	ModuleList map[string]([]*Module)
}

func InitModule(){
	initPlatform()
	platformCount := platformManager.PlatformList.Size()
	moduleManager = ModuleManager{ModuleList: make(map[string][]*Module, platformCount)}
	loadModule()
	logger.Logger.Trace("InitModule success")
}

func loadModule(){
	for _, platform := range platformManager.PlatformList{
		pform,_ := platform.(string)
		filePath := GetPlatformPath(pform)
		fmt.Println(filePath);
		list := make([]*Module,0)
		if conf ,_ := util.ReadFile(filePath); len(conf) > 0 {
			json.Unmarshal(conf, &list)
			fmt.Println(string(conf));
		}
		moduleManager.ModuleList[pform] = list
	}
	str,_ := json.Marshal(moduleManager.ModuleList)
	logger.Logger.Trace(string(str))
}

func AddModule(platform, module string) bool{
	if !ExistPlatform(platform) {
		AddPlatform(platform)
		moduleManager.ModuleList[platform] = make([]*Module,0)
	}
	if !ExistModule(platform,module){
		item := &Module{Platform:platform, Name:module, Queue:container.NewCQueue()}
		moduleManager.ModuleList[platform] = append(moduleManager.ModuleList[platform],item)
		return SaveModule(platform, moduleManager.ModuleList[platform])
	}
	return true
}

func ExistModule(platform, moduleName string) bool{
	if moduleList,isOK := moduleManager.ModuleList[platform]; isOK{
		for _, module := range moduleList{
			if module.Name == moduleName {
				return true
			}
		}
	}
	return false
}

func GetModule(platform, moduleName string) *Module{
	if moduleList,isOK := moduleManager.ModuleList[platform]; isOK{
		for _, module := range moduleList{
			if module.Name == moduleName {
				return module
			}
		}
	}
	return nil
}

func DeleteModule(platform, moduleName string) bool{
	if !ExistPlatform(platform) {
		return false
	}
	if moduleList,isOK := moduleManager.ModuleList[platform]; isOK{
		for index, module := range moduleList{
			if module.Name == moduleName {
				moduleManager.ModuleList[platform] = append(moduleList[:index], moduleList[index+1:]...)
				return SaveModule(platform, moduleManager.ModuleList[platform])
			}
		}
	}
	return false
}

func SaveModule(platform string, moduleList []*Module) bool{
	if jsonStr,ok := json.Marshal(moduleList); ok == nil{
		return SavePlatform(platform, jsonStr)
	}
	return false
}

func GetModuleList() map[string]([]*Module){
	return moduleManager.ModuleList
}

func AddMessageQueue(msg *Message) bool{
	moduleList := moduleManager.ModuleList[msg.Platform]
	for _, module := range moduleList{
		if module.Name == msg.Module {
			logger.Logger.Tracef("AddMessageQueue url=%s", msg.Url)
			module.Queue.Enqueue(msg);
			return true
		}
	}
	return false
}

func TestModule(){
	AddModule("orcp", "module1")
	AddModule("orcp", "module2")
	AddModule("mod", "module1")
	AddModule("orcp", "module3")
	AddModule("orp", "module31")
	AddModule("orp", "module311")
	AddModule("orp", "module3")
	DeleteModule("mod", "module1")
	module := GetModule("mod", "module1")
	jsonStr,_ := json.Marshal(module)
	fmt.Println("GetModule:" + string(jsonStr))
}