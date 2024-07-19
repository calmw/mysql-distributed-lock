package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

func InitConfig() {
	log.Println("init config")
	if _, err := toml.DecodeFile("./config.toml", &Config); err != nil {
		panic("read toml file err: " + err.Error())
		return
	}
}
