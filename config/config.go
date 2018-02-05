package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Dev bool `yaml:"dev"`
	Port string `yaml:"port"`
}

var conf *Config

func parse (filename string, conf *Config) error {
	var err error
	var buf []byte

	buf, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	return err
}

func GetConfig (filename string) *Config {
	if conf == nil {
		conf = new(Config)
		parse(filename, conf)
	}
	return conf
}