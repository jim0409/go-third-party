package plugin

import (
	"log"
)

func init() {
	log.Println("plugin1 init")
}

func Run() {
	log.Println("run plugin!!")
}
