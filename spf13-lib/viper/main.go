package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
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
	cfg := &Configs{}
	cfg.loadConfig(configPath)
	for k, v := range cfg.config.AllKeys() {
		fmt.Println(k, v)
	}

}
