package package1

import (
	"fmt"

	"go-third-party/tomlAst/plugins"
)

type DemoStruct struct {
	Test int `toml:"test"`
}

func (d *DemoStruct) Gather() error {
	fmt.Println("Enther here XDDDDD")
	fmt.Println(d.Test)
	return nil
}

func init() {
	plugins.Add("package1", func() plugins.Plugin {
		return &DemoStruct{}
	})
}
