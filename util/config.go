package util

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Config struct{

}

const(
	PlatformConf = "./conf/platform.conf"
	MessageConf = "./conf/message.conf"
)
var LmqConfig Config

func LoadConfig(configFile string) error {
	if _, err := toml.DecodeFile(configFile, &LmqConfig); err != nil {
		return  err
	}
	return nil
}

func LoadJson(jsonFile string) []byte{
	if byte, err := ioutil.ReadFile(jsonFile); err == nil{
		return byte
	}
	return nil
}


