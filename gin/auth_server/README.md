# intro
1. 需要以golang的gin框架，創造一個驗證使用者的JWT的api... AuthCheck()
2. 需要對某些情境下的API做額外驗證... CheckParamAndHeader(h gin.HandlerFunc) gin.HandlerFunc

### notes:
1. c.JSON(`status code`, gin.H{ map[string]interface{} }): 回傳一個完整的rest JSON格式給client端
2. c.Abort(): 放棄這次文本，強制返回

# refer:
gin中使用Route
- https://www.jianshu.com/p/d4b52187d233
httpServer中使用graceful shutdown
- https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97
gin中使用裝飾器Decorator
- https://www.jianshu.com/p/55d65dd748ca
gin中使用middleware
- http://www.ttlsa.com/golang/gin-middleware-example/
golang中使用io.ReadWriter拿取Request body而不會失去body的內容
- https://stackoverflow.com/questions/32008680/gin-go-lang-how-to-use-context-request-body-and-retain-it
- https://medium.com/@xoen/golang-read-from-an-io-readwriter-without-loosing-its-content-2c6911805361
- https://stackoverflow.com/questions/47186741/how-to-get-the-json-from-the-body-of-a-request-on-go/47295689#47295689

### extend-refer:
一般的http handler中使用Decorator
- https://colobu.com/2019/08/21/decorator-pattern-pipeline-pattern-and-go-web-middlewares/


### 一般資料結構下的Decorator
一般的decorator是透過宣告`Wrapper`返回重複的`Component接口`。並以`w.ReturnInt()`或`w.ReturnString()`做驗證
```go
package decorator

type Component interface {
	ReturnInt() int
	ReturnString() string
}

type Validator struct {
	Component
}

func WrapValidator(c Component) Component {
	return &Validator{
		Component: c,
	}
}

func (w *Validator) ReturnInt() int {
	if w.Component.ReturnInt() >= 20 {
		if w.Component.ReturnInt() >= 30 {
			return w.Component.ReturnInt() - 10
		} else {
			return w.Component.ReturnInt() + 1
		}
	}
	return w.Component.ReturnInt()
}

func (w *Validator) ReturnString() string {
	return w.Component.ReturnString() + " test"
}
```