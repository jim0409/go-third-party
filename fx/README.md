# intro

Fx 是一個 golang 版本的依賴注入框架，透過可重用、可組合的模塊化擴建`模組`

# Fx 生命週期

Fx 生命週期

### 1. App.New
1. 新建 App
  - 指定 container
  - 指定 lifecycle
  - 指定 startTimeout
  - 指定 stopTimeout
  - 指定 logger

2. 應用 Option
  - errorHookOption
  - invokeOption
  - provideOption

3. 提供注入對象的 struct{}
<!-- 注意 provide: LifeCycle, shutdowner, dotGraph 三個特殊處理 -->

4. 執行 invoke 提供的方法(所有的構造函數都是延遲性執行，invoke提供的函數能夠立刻執行)


### 2. App.Start
1. 通過使用 withTimeout 來完成 app 啟動: 使用 context.Context 紀錄本次操作結果
>  調用 app.start 方法來執行啟動。根據屬性 startTimeout 來有效控制 app 啟動時間


### 3. Inject.reflect
1. 此處需要根據 New 過程中設置的 option 來進行操作，獲取注入對象名關連的具體對象


### 4. App.Stop
通過使用 withTimeout 來完成 app 停止:
使用 context.Context 紀錄本次操作結果
調用 app.lifetcycle.stop 方法來執行關閉
根據屬性 stopTimeout 來有效控制 app 關閉時間


# Fx 應用
1. 一般步驟
- 按需定義一些構造函數: 主要用於生成依賴注入的具體類型
```go
type FxDemo struct {
	// 字段式可導出的，是由於 golang 的 reflection 的特性決定: 必須能夠導出且可尋址才能進行設置
	Name string
}
```

```go
func NewFxDemo() {
	return FxDemo {
		Name: "hello, world",
	}
}
```
- 使用 Provide 將具體反射的類型添加到 container 中，可以按需添加任意多個函數
> fx.Provide(NewFxDemo)

- 使用 Populate 完成變量與具體類型的映射
```go
var fx FxDemo
fx.Populate(fx)
```

- 新建 app 對象(application 容器包括定義注入變量、類型、不同對象 lifecycle 等)
```go
app := fx.New(
	fx.Provide(NewFxDemo,),    // 構造函數可以任意多個
	fx.Populate(new(FxDemo)),  // 反射變量也可以任意多個，並不需要和上面構造函數對應
)

app.Start(context.Backgroud())         // 開啟 container
defer app.Stop(context.Background())   // 關閉 container
```

- 使用
```go
fmt.Printf("the result is %s\n", fx.Name)
```

# 簡單 Demo
將`io.reader`與具體實現類關聯起來
```go
package main

import (
	"context"
	"fmt"
	// "github.com/uber-go/fx"
	"go.uber.org/fx"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var reader io.Reader

	app := fx.New(
		// io.reader 的應用
		fx.Provide(func() io.Reader {
			return strings.NewReader("hello world")
		}),
		fx.Populate(&reader),  // 通過依賴注入完成變量與具體類的映射
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	// 使用，reader變量已與 fx.Provide 注入的實現類關聯了
	bs, err := ioutil.ReadAll(reader)
	if err != nil{
		log.Panic("read occur error, ", err)
	}
	fmt.Printf("the result is '%s' \n", string(bs))
}

// 輸出
/*
[Fx] PROVIDE    io.Reader <= main.main.func1()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUNNING
*/
```


# 使用 Struct 參數
前面的使用方式一旦需要進行注入的類型過多，可以通過`struct`參數方式來解決

```go
package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	type t3 struct {
		Name string
	}

	type t4 struct {
		Age int
	}

	var (
		v1 *t3
		v2 *t4
	)

	app := fx.New(
		fx.Provide(func() *t3 { return &t3{"hello everybody!!!"} }),
		fx.Provide(func() *t4 { return &t4{2019} }),

		fx.Populate(&v1),
		fx.Populate(&v2),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	fmt.Printf("the result is %v, %v\n", v1.Name, v2.Age)
}

// 輸出
/*
[Fx] PROVIDE    *main.t3 <= main.main.func1()
[Fx] PROVIDE    *main.t4 <= main.main.func2()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUNNING
the result is hello everybody!!!, 2019
*/
```

> 但是通過`Provide`提供構造函數是生成相同類型會有多個值的問題

# 使用 struct{} + Name 標籤
在`Fx`未使用`Name`或`Group`標籤時，不允許存在多個相同類型的構造函數，一旦存在會觸發`panic`

```go
package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	type t3 struct {
		Name string
	}

	// name 標籤的使用
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

	fmt.Printf("the result is %v, %v\n", targets.V1.Name, targets.V2.Name)
}

// 輸出
/*
[Fx] PROVIDE    *main.t3[name = "n1"] <= main.main.func1()
[Fx] PROVIDE    *main.t3[name = "n2"] <= main.main.func1()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUNNING
the result is hello-HELLO, world-WORLD
*/
```

> 通過 Name 標籤，可以完成在`Fx`容器注入相同類型

# 使用 struct{} + Group 標籤
使用`group`標籤同樣也能完成上面的功能
```go
package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	type t3 struct {
		Name string
	}

	// 使用 group 標籤
	type result struct {
		fx.Out

		V1 *t3 `group:"g"`
		V2 *t3 `group:"g"`
	}

	targets := struct {
		fx.In

		Group []*t3 `group:"g"`
	}{}

	app := fx.New(
		fx.Provide(func() result {
			return result{
				V1: &t3{"hello-000"},
				V2: &t3{"world-www"},
			}
		}),

		fx.Populate(&targets),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	for _, t := range targets.Group {
		fmt.Printf("the result is %v\n", t.Name)
	}
}


// 輸出
/*
[Fx] PROVIDE    *main.t3[group = "g"] <= main.main.func1()
[Fx] PROVIDE    *main.t3[group = "g"] <= main.main.func1()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUNNING
the result is hello-000
the result is world-www
*/
```


# refer:
- https://blog.csdn.net/h_sn9999/article/details/120524130