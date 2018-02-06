package test

import (
	"testing"
	"goBuy/config"
	"goBuy/db"
)

type Users struct {
	Id int
	Name string
}

func TestConn(t *testing.T) {
	config.Load("config_test.yaml")
	db.Connect(config.Conf.Mysql)
	user := Users{}
	db.MysqlConns["default"].First(&user)
	if user.Name != "test" {
		t.Error("Connect error")
	}
}
