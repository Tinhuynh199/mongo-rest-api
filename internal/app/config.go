package app

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Mongo  MongoConfig  `yaml:"mongo"`
}

type ServerConfig struct {
	Name string `yaml:"name"`
	Port *int64 `yaml:"port"`
}

type MongoConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
}

func (c *Config) getConf() *Config {

	yamlFile, err := ioutil.ReadFile("./configs/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func GetConfig() *Config {
	var conn Config
	conn.getConf()
	server := conn.Server
	mongo := conn.Mongo

	return &Config{
		Server: server,
		Mongo: mongo,
	}
}
