# intro
arangodb is a no-sql db with fast insert speed and mult-models properties


# quick-start
setup a arangodb single node with `docker-compose` (note with default usr/password : root/password for `arangodb login`)
> docker-compose up -d


# start a cluster arangodb
> refer: https://github.com/jim0409/LinuxIssue/tree/master/arangodb-cluster

# how to execute

> go run main.go --name root --password --addr http://127.0.0.1:8529


# refer:
how to use client sdk
- https://github.com/arangodb/go-driver/tree/master

how to deploy cluster env
- https://gist.github.com/neunhoef/1620c6c50e84a12be2b476bed419c644
