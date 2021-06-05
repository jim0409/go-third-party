# quick start
1. 啟動docker-compose帶起對應的kafka以及zookeeper
> docker-compose up -d

2. 啟動producer.go往已經建立好的topic(sarama)送資料
> go run producer.go localhost:9092 sarama

3. 啟動consumer.go來消費線上的kafka topic
> go run consumer.go localhost:9092 1 sarama

# refer:
- https://github.com/Shopify/sarama
- https://juejin.im/post/5d40f179f265da038f47e9eb


### kafka 工作原理
- http://xstarcd.github.io/wiki/Cloud/kafka_Working_Principles.html
- https://www.zhihu.com/question/28925721


### 使用kafka binary 實作
- https://medium.com/%E3%84%9A%E5%8C%97%E7%9A%84%E6%89%80%E8%A6%8B%E6%89%80%E8%81%9E/kafka-%E8%A8%AD%E8%A8%88%E6%80%9D%E8%80%83-partition-replication-adc9eac58e36


# kafka-docker-refer:
- https://github.com/wurstmeister/kafka-docker


 
# Automatically create topics
If you want to have kafka-docker automatically create topics in Kafka during creation, a 
KAFKA_CREATE_TOPICS environment variable can be added in docker-compose.yml.

Here is an example snippet from docker-compose.yml:
```sh
    environment:
      KAFKA_CREATE_TOPICS: "Topic1:1:3,Topic2:1:1:compact"
```

Topic 1 will have 1 partition and 3 replicas,
Topic 2 will have 1 partition, 1 replica and a cleanup.policy set to compact. 
