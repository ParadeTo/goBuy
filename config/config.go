package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type MysqlConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname string `yaml:"dbname"`
}

type Config struct {
	Dev bool `yaml:"dev"`
	Port string `yaml:"port"`
	Mysql []MysqlConfig `yaml:"mysql"`
}

var Conf *Config

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

// singleton pattern
func Load (filename ...string) {
	defaultFilename := "config.yml"

	if filename != nil {
		defaultFilename = filename[0]
	}

	if Conf == nil {
		Conf = new(Config)
		parse(defaultFilename, Conf)
	}
}