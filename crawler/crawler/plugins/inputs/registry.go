package inputs

import "github.com/jimweng/crawler/crawler/utils"

type Creator func() utils.Input

var Inputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Inputs[name] = creator
}
