package util

import (
	"github.com/BurntSushi/toml"
	"lmq/util/logger"
)

const(
	LmqConf = "./conf/lmq.conf"
)

type Config struct{
	Serve struct{
		HttpPort int
	}
}

var LmqConfig Config

func LoadConfig() error {
	if _, err := toml.DecodeFile(LmqConf, &LmqConfig); err != nil {
		return  err
	}
	return nil
}

func ConfigTest(){
	if isok := LoadConfig();isok != nil{
		logger.Logger.Tracef("loadConfig err=%s", isok)
	}
	logger.Logger.Tracef("httpPort:%d",LmqConfig.Serve.HttpPort);
}

