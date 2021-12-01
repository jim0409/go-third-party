# Fx 定義

1. Annotated(位於 annotated.go 文件)主要用於採用`annotated`的方式，提供`Provide`注入類型
```go
type t3 struct {
	Name string
}

targets := struct {
	fx.In

	V1 *t3 `name:"n1"`
}{}

app := fx.New(
	fx.Provide(fx.Annotated{
		Name: "n1",
		Target: func() *t3{
			return &t3{"hello world"}
		},
	}),
	fx.Populate(&targets)
)
app.Start(context.Background())
defer app.Stop(context.Background())

fmt.Printf("the result is = '%v'\n", targets.V1.Name)
// 源碼中`Name`和`Group`兩個字段與前面提到的`Name`標籤和`Group`標籤是一樣的，只能選其一使用
```


2. App(位於app.go文件)，提供注入對象具體的容器、LiftCycle、容器的啟動及停止、類型變量及實現類注入和兩者映射等操作
```go
type App struct {
	err 		error
	container	*dig.Container		// 容器
	lifecycle	*lifecycleWrapper	// 生命週期
	provides	[]interface{}		// 注入的類型實現類
	invokes		[]interface{}
	logger		*fxLog.Logger
	startTimeout	time.Duratino
	stopTimeout	time.Duratino
	errorHooks	[]ErrorHandler

	donesMu		sync.RWMutex
	dones		[]chan os.Signal
}

// 新建一個 App 對象
func New(opts ...Option) *App {
	logger := fxlog.New()	// 紀錄 Fx 日誌
	app := &App{
		container:	dig.New(dig.DeferAcyclicVerification()),
		lifecycle:	lc,
		logger:		logger,
		startTimeout:	DefaultTimeout,
		stopTimeout:	DefaultTimeout,
	}

	for _, opt := range opts { // 提供的 Provide 和 Populate 的操作
		opt.apply(app)
	}

	// 進行 app 相關一些操作
	for _, p := range app.provides {
		app.provide(p)
	}

	app.provide(func() Lifecycle { return app.lifecycle })
	app.provide(app.shutdowner)
	app.provide(app.dotGraph)

	if app.err != nil { // 紀錄 app 初始化過程是否正常
		app.logger.Printf("Error after options were applied: %v", app.err)
		return app
	}

	// 執行 invoke
	if err := app.executeInvokes(); err != nil{
		app.err = err

		if dig.CanVisualizeError(err) {
			var b bytes.Buffer
			dig.Visualize(applcontainer, &b, dig.VisualizeError(err))
			err = errorWithGrap{
				graph:	b.String(),
				err:	err,
			}
		}
		errorHandlerList(app.errorHooks).HandleError(err)
	}
	return app
}
```

3. Extract(位於extract.go文件)
主要用於在 application 啟動初始化過程，通過依賴注入的方式將容器中的變量值填充給定的 struct

其中 target 必須是指向 struct 的指針，並且只能填充可導出的字段

(golang只能通過反射修改可導出並且可尋址的字段)

4. 其他
諸如 Populate 是用來替換 Extract 的，而 LiftCycle 和 inout.go 涉及內容較多

# 其他

在 Fx 中提供的構造函數都是惰性調用，可以通過`invocations`在`application`啟動來完成一些必要的初始化工作: fx.Invoke(function);

也可以按需自定義實現`LifeCycle`的`Hook`對應的`OnStart`和`OnStop`用來完成手動啟動容器和關閉，來滿足一些自己實際的業務需求

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
)

// Logger 構造函數
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags*/)
	logger.Print("Executing NewLogger")
	return logger
}


// http.Handler 構造函數
func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// http.ServeMux 構造函數
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux")

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:		":8080",
		Handler:	mux,
	}

	lc.Append(fx.Hook{	// 自定義生命週期過程對應的啟動和關閉行為
		OnStart: func(ctx context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// 註冊 http.Handler
func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewHandler,
			NewMux,
		),
		fx.Invoke(Register),	// 通過 invoke 來完成 Logger、Handler、ServeMux 的創建
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil { // 手動調用 Start
		log.Fatal(err)
	}

	http.Get("http://localhost:8080/")	// 具體操作

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil { // 手動調用 Stop
		log.Fatal(err)
	}
}
```

# Fx 源碼解析
主要包括 app.go、lifecycle.go、annotated.go、populate.go、inout.go、shutdown.go、extract.go(可以忽略，了解populate.go)

以及輔助的 internal 中的 fxlog、fxreflect、lifecycle


# refer:
- https://blog.csdn.net/h_sn9999/article/details/120524130