package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, domain, obj, act string) {
	ok, _ := e.Enforce(sub, domain, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj, domain)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj, domain)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "dajun", "tenant1", "data1", "read") // true
	check(e, "dajun", "tenant2", "data2", "read") // true

	/*
		在 tenant1 中，只有 admin 可以讀取數據 data1
		在 tenant2 中，只有 admin 可以讀取數據 data2

		dajun 在 tenant1 中是 admin，但在 tenant2 中不是
	*/
	check(e, "dajun", "tenant1", "data1", "write") // true
	check(e, "dajun", "tenant2", "data2", "write") // false
}
