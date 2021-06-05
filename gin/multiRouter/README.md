# intro
使用多個group來做複數的router

# quick start
1. test hello(v1/v2)
> curl localhost:8080/v1/hello
```
[GIN] 2019/10/30 - 14:12:24 | 200 |      41.168µs |             ::1 | GET      /v1/hello
```

2. test goodbye(v1/v2)
> curl -XPOST localhost:8080/v1/goodbye
```
[GIN] 2019/10/30 - 14:12:43 | 200 |      22.327µs |             ::1 | POST     /v1/goodbye
```

# refer:
- https://stackoverflow.com/questions/42373423/how-to-add-multiple-groups-to-gin-routing-for-api-version-inheritance