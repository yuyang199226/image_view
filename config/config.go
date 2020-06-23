package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
	"time"
	//   "github.com/sirupsen/logrus"
)

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
	Debug   bool
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

var (
	cfg   *tomlConfig
	once  sync.Once
	mutex sync.Mutex
)

func Config() *tomlConfig {
	once.Do(ReloadConfig)
	//logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true, FullTimestamp: true, TimestampFormat: "02-01-2006 15:04:05.000000"})
	return cfg

}

func ReloadConfig() {
	filePath, err := filepath.Abs("config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse toml filePath %s\n", filePath)
	config := new(tomlConfig)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}
	mutex.Lock()
	defer mutex.Unlock()
	cfg = config
}
