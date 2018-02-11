package db

import (
	"testing"
	"goBuy/config"
	"goBuy/db"
	"github.com/stretchr/testify/assert"
)

type Users struct {
	Id int
	Name string
}

func TestConn(t *testing.T) {
	config.Load("config_test.yml")
	db.Connect(config.Conf.Mysql)
	user := Users{}
	db.MysqlConns["default"].First(&user)
	assert.Equal(t, user.Name, "test")
}
