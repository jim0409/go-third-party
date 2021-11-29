package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

/*
通過 Name 標籤可以完成在 Fx 容器注入相同類型
*/

func main() {
	type t3 struct {
		Name string
	}

	type result struct {
		fx.Out

		V1 *t3 `name:"n1"`
		V2 *t3 `name:"n2"`
	}

	targets := struct {
		fx.In

		V1 *t3 `name:"n1"`
		V2 *t3 `name:"n2"`
	}{}

	app := fx.New(
		fx.Provide(func() result {
			return result{
				V1: &t3{"hello-HELLO"},
				V2: &t3{"world-WORLD"},
			}
		}),
		fx.Populate(&targets),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	fmt.Printf("the result is %v, %v \n", targets.V1.Name, targets.V2.Name)
}
