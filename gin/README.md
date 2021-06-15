# intro
簡單使用gin server用來做測試`http`及`ws`
```log
[GIN-debug] GET    /ping                     --> main.setupRouter.func1 (3 handlers)
[GIN-debug] POST   /test                     --> main.setupRouter.func2 (3 handlers)
[GIN-debug] GET    /test                     --> main.setupRouter.func3 (3 handlers)
[GIN-debug] PUT    /test                     --> main.setupRouter.func4 (3 handlers)
[GIN-debug] DELETE /test                     --> main.setupRouter.func5 (3 handlers)
[GIN-debug] GET    /                         --> main.setupRouter.func6 (3 handlers)
[GIN-debug] GET    /ws                       --> main.wsHandler (3 handlers)
```


# refer
- https://github.com/gin-gonic/gin