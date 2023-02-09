package main

import (
	ini "gopkg.in/ini.v1"
)

type Config struct {
	BaseConf
	SmtpConf
}

type BaseConf struct {
	Env string `ini:"Env"`
}

type SmtpConf struct {
	Auth   string `ini:"Auth"`
	Host   string `ini:"Host"`
	User   string `ini:"User"`
	Server string `ini:"Server"`
}

func InitConfig(confPath string) (*Config, error) {
	conf := new(Config)
	if err := ini.MapTo(conf, confPath); err != nil {
		return nil, err
	}

	return conf, nil
}
