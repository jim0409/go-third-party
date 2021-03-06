package outputs

import "go-third-party/crawler/crawler/utils"

type Creator func() utils.Output

var Outputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Outputs[name] = creator
}
