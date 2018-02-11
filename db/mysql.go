package db

import (
	"goBuy/config"
	. "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var MysqlConns = make(map[string]*DB)

// Usage:
// 	config.Load("config_test.yml")
//  db.Connect(config.Conf.Mysql)
//  db.MysqlConns["default"].First(&user)
func Connect(mysqlConfigs []config.MysqlConfig) {
	for _, mysql := range mysqlConfigs {
		params := mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Dbname + "?charset=utf8&parseTime=True&loc=Local"
		conn, err := Open("mysql", params)
		if err != nil {
			log.Fatalln(err)
			return
		}
		conn.DB().SetMaxIdleConns(10)
		conn.DB().SetMaxOpenConns(100)
		MysqlConns[mysql.Name] = conn
	}
}