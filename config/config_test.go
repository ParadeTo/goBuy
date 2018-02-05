package config

import (
	"testing"
	"os"
)

func TestConfig(t *testing.T) {
	yamlcontext := `dev: true
port: 8080	`
	f, err := os.Create("testconfig.yaml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(yamlcontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("testconfig.yaml")

	conf := GetConfig("testconfig.yaml")
	if !conf.Dev {
		t.Error("parse error")
	}
}
