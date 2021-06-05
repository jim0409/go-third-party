# This is a crawler system
1. use crawler to parse website data into mysql
2. with grpc_server and grpc_client throw data into front end
3. architecture as below;
```
crawler --> mysql

          |-- grpc_server -- grpc_client(net/http listen server) -- |
mysql ----|                                                         |-- nginx
          |-- grpc_server -- grpc_client(net/http listen server) -- |
```
4. query with specific `pkg name`
> curl http://$URI/query -d '{"data":$pkg_name}'
> e.g. curl http://localhost/query -d '{"data":"cgi"}'

# Requirement
1. `git` (version: 2.17.1 or later)
2. `docker` (version: 18.09.7)
3. `docker-compose`

# Installation
- With git cli
> git clone https://github.com/jim0409/jimweng.git
- Move to this folder
> cd ./crawler

# Command to build the environment
Under this directory, execute command below
> docker-compose up --scale grpcclient=2 --scale grpcserver=2 -d

# note
### Future Work
1. parallism crawler components
2. refactor grpc client and server
3. add `health check`, `prometheus` and `monitor` for better analysis
4. use `decorator mode`/`redis` to cache query result in grpc client

### Known Issue
1. crawler components would use `gorm.Open` everytime, could separate the function to extra interface
> workaround, `set idle connections as 0 to make sure no more connection while crawler is not working`

# refer
- way to use docker-compose with network
https://docs.docker.com/compose/compose-file/#network_mode

- way to use gorm as ORM in golang
https://gorm.io/docs/

- way to use grpc
https://developers.google.com/protocol-buffers/docs/gotutorial

- way to modify nginx config
http://nginx.org/en/docs/http/load_balancing.html