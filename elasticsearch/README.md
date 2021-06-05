# quick start 
1. build environment 
> docker-compose up -d

2. run porcess
> go run main.go

3. execute `curl`
> curl -XPOST "http://localhost:8080/documents" -d @document.json

4. check data with end points
> http://127.0.0.1:9200/documents/_search?pretty

or with (kibana) browser
> http://127.0.0.1:5601

# refer:
how to startup a elasticsearch
- https://github.com/tinrab/go-elasticsearch-example

knowledge of gin with `BindJSON`
- https://blog.csdn.net/heart66_A/article/details/100796964
