package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) (bool, string) {
	var msg string
	ok, _ := e.Enforce(sub, obj, act)

	if ok {
		msg = fmt.Sprintf("%s CAN %s %s\n", sub, act, obj)
	} else {
		msg = fmt.Sprintf("%s CANNOT %s %s\n", sub, act, obj)
	}
	return ok, msg
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
	check(e, "root", "data2", "read")
}
