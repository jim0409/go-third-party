package main

import (
	"fmt"

	arango "github.com/adamwasila/arangodb-adapter"
	casbin "github.com/casbin/casbin/v2"
)

func main() {
	a, err := arango.NewAdapter(
		arango.OpBasicAuthCredentials("root", ""),
		// arango.OpAutocreate(false),
		arango.OpAutocreate(true), // 自動創建 collection
		arango.OpDatabaseName("_system"),
		arango.OpEndpoints("http://127.0.0.1:8529"),
		arango.OpCollectionName("PolicyRules"),
		arango.OpFieldMapping("p", "sub", "obj", "act"),
	)

	if err != nil {
		fmt.Printf("Adapter creation error! %s\n", err)
		return
	}

	e, err := casbin.NewSyncedEnforcer("model.conf", a)
	if err != nil {
		fmt.Printf("Enforcer creation error! %s\n", err)
		return
	}
	err = e.LoadPolicy()
	if err != nil {
		fmt.Printf("Load policy error! %s\n", err)
		return
	}

	injectPolicy(e)

	sub, obj, act := "boss", "owner", "book"
	fmt.Println(checkPolicy(e, sub, obj, act))
	sub, obj, act = "boss", "writer", "book"
	fmt.Println(checkPolicy(e, sub, obj, act))
	sub, obj, act = "boss", "reader", "book"
	fmt.Println(checkPolicy(e, sub, obj, act))

}

var (
	groups   = []string{"admin", "developer", "customer"}
	roles    = []string{"owner", "writer", "reader"}
	policies = []string{"business", "market", "tutorial", "book"}
)

func checkPolicy(e *casbin.SyncedEnforcer, sub, obj, act string) string {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		return fmt.Sprintf("Failed to enforce! %s", err)

	}

	if !ok {
		return fmt.Sprintf("%s %s %s: Forbidden!", sub, obj, act)
	}

	return fmt.Sprintf("%s %s %s: Access granted", sub, obj, act)
}

func injectPolicy(e *casbin.SyncedEnforcer) {

	e.AddRoleForUser("boss", groups[0])
	for i := 0; i < len(policies); i++ {
		// 保證 admin 可以同時使用 owner/ writer/ reader，否則只要加入roles[0]即可
		e.AddPolicy(groups[0], roles[0], policies[i])
		e.AddPolicy(groups[0], roles[1], policies[i])
		e.AddPolicy(groups[0], roles[2], policies[i])
	}

	e.AddRoleForUser("employee", groups[1])
	for i := 1; i < len(policies); i++ {
		// 同 admin 原因
		e.AddPolicy(groups[1], roles[1], policies[i])
		e.AddPolicy(groups[1], roles[2], policies[i])
	}

	e.AddRoleForUser("villager", groups[2])
	e.AddRoleForUser(groups[2], roles[2])
	for i := 2; i < len(policies); i++ {
		e.AddPolicy(groups[2], roles[2], policies[i])
	}

	e.SavePolicy()

	fmt.Printf("Thats all folks\n")
}
