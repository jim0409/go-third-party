package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	// check(e, "dajun", "data", "read")
	// check(e, "dajun", "data", "write")
	// check(e, "lizi", "data", "read")
	// check(e, "lizi", "data", "write")
	check(e, "dajun", "prod.data", "read")
	check(e, "dajun", "prod.data", "write")
	check(e, "lizi", "dev.data", "read")
	check(e, "lizi", "dev.data", "write")
	check(e, "lizi", "prod.data", "write")
	/*
		dajun 屬於 admin 角色
		lizi 屬於 developer 角色
		prod.data 屬於生產資源 `prod` 角色
		dev.data 屬於開發資源 `dev` 角色

		admin 角色擁有對 `prod` 和 `dev` 類資源的 讀寫 權限
		developer 只能對 `dev` 的讀寫權限 和 `prod` 的 讀權限
	*/

}
