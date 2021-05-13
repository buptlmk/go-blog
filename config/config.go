package config

import (
	"encoding/json"
	"io/ioutil"
)

type SettingsStruct struct {
	ProjectPath string
	Log         string
	SecretKey   string
	SQLSet      SQLSettings
	Server      ServerSettings
	RedisSet    RedisSettings
	RabbitMq    RabbitMqSet
}

var Settings SettingsStruct

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Settings)
	if err != nil {
		panic(err)
	}
}
