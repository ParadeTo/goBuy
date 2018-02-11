package config

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	Load("config_test.yml")

	fmt.Println(Conf)
	assert.Equal(t, Conf.Dev, true)
	assert.Equal(t, Conf.Mysql[0].Name, "default")
}
