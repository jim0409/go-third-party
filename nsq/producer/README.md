# intro - Producer Method
- `NewProducer`  would generate an new producer
- `Publis` would publish msg to queue manager system
- `Stop` would stop the connection within system

# Way to produce some msg to Nsq
> curl -XPOST localhost:4151/pub?topic=TopicName -d '{"msgkey": "msgValue"}'

# refer:
- https://blog.csdn.net/tian_lai_yuyuh/article/details/52700695