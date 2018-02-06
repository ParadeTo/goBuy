package test

import (
	"testing"
	. "goBuy/config"
	//"fmt"
	"fmt"
)

func TestConfig(t *testing.T) {
	Load("config_test.yaml")

	fmt.Println(Conf)
	if !Conf.Dev {
		t.Error("parse error")
	}

	if Conf.Mysql[0].Name != "default" {
		t.Error("parse error")
	}
}
