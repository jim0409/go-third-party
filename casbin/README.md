# intro
`casbin`將訪問控制模型抽象到一個基於`PERM(Policy, Effect, Request, Matchers`
元哦模型的配置文件(模型文件)中。

因此切換或更新授權機制只要簡單的修改配置文件

- `policy` 是策略或者說是規則的定義。他定義了具體的規則
- `request` 是對訪問請求的抽象，他與`e.Enforce()`函數的參數是一對一對應的
- `matcher` 匹配器會將請求與定義的每個`policy`一一匹配，生成多個匹配結果
- `effect` 根據對請求運用匹配器得出的所有結果進行匯總，來決定該請求是`允許`還是`拒絕`

# UML
```
                                 Policy 1
                                   |
                                   v
                            ---> Matcher -- Policy 1 ----
                            |     exp        Effect     |
                            |                           |
                            |    Policy 2               |
                            |      |                    |    --------------
                            |      v                    ---> | Effect     |
Request(sub, verb, action) --- > Matcher -- Policy 2 ------> | Expression | --> Resulte (true/ false)
                            |     exp        Effect     ---> |            |
                            |                           |    --------------
                            |    Policy 3               |
                            |      |                    |
                            |      v                    |
                            ---> Matcher -- Policy 3 ----
                                  exp        Effect
```

# refer:
- https://juejin.cn/post/6844904191257739277
