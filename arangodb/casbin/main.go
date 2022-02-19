package main

import (
	"fmt"

	arango "github.com/adamwasila/arangodb-adapter"
	casbin "github.com/casbin/casbin/v2"
)

func main() {
	a, err := arango.NewAdapter(
		arango.OpBasicAuthCredentials("root", ""),
		arango.OpAutocreate(false),
		arango.OpDatabaseName("Database"),
		arango.OpEndpoints("http://127.0.0.1:8529"),
		arango.OpCollectionName("PolicyRules"),
		arango.OpFieldMapping("p", "sub", "obj", "act"),
	)

	if err != nil {
		fmt.Printf("Adapter creation error! %s\n", err)
		return
	}

	e, err := casbin.NewEnforcer("model.conf", a)
	if err != nil {
		fmt.Printf("Enforcer creation error! %s\n", err)
		return
	}
	err = e.LoadPolicy()
	if err != nil {
		fmt.Printf("Load policy error! %s\n", err)
		return
	}

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	e.AddPolicy("adam", "data1", "write")
	e.AddPolicy("bob", "data1", "read")
	e.AddPolicy("cecile", "data1", "write")
	e.AddPolicy("alice", "data1", "read")
	e.SavePolicy()

	fmt.Printf("Thats all folks\n")

	sub, obj, act := "alice", "data1", "read"
	r, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Printf("Failed to enforce! %s\n", err)
		return
	}
	if !r {
		fmt.Printf("%s %s %s: Forbidden!", sub, obj, act)
	} else {
		fmt.Printf("%s %s %s: Access granted", sub, obj, act)
	}
}
