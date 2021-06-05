package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	logRoot     = "service.log"
	mysqlRoot   = "service.mysql"
	serviceRoot = "service.service"
)

type Configs struct {
	config *viper.Viper
}

func (c *Configs) loadConfig(path string) {
	c.config = viper.New()

	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	c.config.SetConfigFile(abs)
	if err := c.config.ReadInConfig(); err != nil {
		panic(err)
	}

	c.config.AllSettings()
}

func main() {
	configPath := "./config.yaml"

	Config := &Configs{}
	Config.loadConfig(configPath)

	fmt.Println(Config)
}
