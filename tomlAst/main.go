package main

import (
	"errors"
	"io/ioutil"

	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
	"go-third-party/tomlAst/plugins"
	_ "go-third-party/tomlAst/plugins/all"
)

type Config struct {
	plugins map[string]*ast.Table
}

func (c *Config) Plugins() map[string]*ast.Table {
	return c.plugins
}

func (c *Config) Apply(name string, v interface{}) error {
	if tbl, ok := c.plugins[name]; ok {
		return toml.UnmarshalTable(tbl, v)
	}

	return nil
}

var ErrInvalidConfig = errors.New("invalid configuration")

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tbl, err := toml.Parse(data)
	if err != nil {
		return nil, err
	}

	c := &Config{
		plugins: make(map[string]*ast.Table),
	}

	for name, val := range tbl.Fields {
		subtbl, ok := val.(*ast.Table)
		if !ok {
			return nil, ErrInvalidConfig
		}

		c.plugins[name] = subtbl

	}

	return c, nil
}

func main() {
	config, err := LoadConfig("./.env.conf")
	if err != nil {
		panic(err)
	}

	ag, err := NewAgent(config)
	if err != nil {
		panic(err)
	}
	if _, err = ag.LoadPlugins(); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	err = ag.crank()
	if err != nil {
		panic(err)
	}

}

type Agent struct {
	TestAgent int
	Config    *Config
	plugins   []plugins.Plugin
}

func NewAgent(config *Config) (*Agent, error) {
	agent := &Agent{Config: config}
	// 當要加入對agent的設定時，可以透過這邊帶入
	err := config.Apply("agent", agent)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (a *Agent) LoadPlugins() ([]string, error) {
	var names []string
	var err error

	for name, creator := range plugins.Plugins {
		plugin := creator()

		err = a.Config.Apply(name, plugin)
		if err != nil {
			return nil, err
		}

		a.plugins = append(a.plugins, plugin)
		names = append(names, name)
	}

	return names, nil
}

func (a *Agent) crank() error {
	for _, plugin := range a.plugins {
		err := plugin.Gather()
		if err != nil {
			return err
		}
	}
	return nil
}
