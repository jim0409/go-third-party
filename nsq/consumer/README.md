# intro - Consumer Method
- `NewConsumer` would generate an new consumer
- `AddHandler` would define the steps after receive msg from queue system
- `ConnectToNSQD` would connect to system
- `Wait` would ask process to wait msg

# Way to produce some msg to Nsq
> curl -XPOST localhost:4151/pub?topic=TopicName -d '{"msgkey": "msgValue"}'

# refer:
- https://blog.csdn.net/tian_lai_yuyuh/article/details/52700695

# way to build a linux binary under mac env
`GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags  " -s -w" -o consumer`
