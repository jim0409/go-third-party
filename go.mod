module go-third-party

go 1.16

require (
	github.com/Jeffail/gabs v1.4.0
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/Shopify/sarama v1.29.1
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/arangodb/go-driver v0.0.0-20210602123439-e63cef3dc348
	github.com/aws/aws-lambda-go v1.24.0
	github.com/bluele/slack v0.0.0-20180528010058-b4b4d354a079
	github.com/casbin/casbin/v2 v2.31.3
	github.com/casbin/gorm-adapter/v3 v3.3.0
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/fatih/structs v1.1.0
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.7.2
	github.com/go-ini/ini v1.62.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.8.4
	github.com/google/gopacket v1.1.19
	github.com/googollee/go-socket.io v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/influxdata/influxdb v1.9.1
	github.com/jinzhu/gorm v1.9.16
	github.com/klauspost/compress v1.13.1 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/minio/minio-go/v7 v7.0.11
	github.com/mitchellh/mapstructure v1.2.2
	github.com/montanaflynn/stats v0.6.6
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1
	github.com/nats-io/nats-streaming-server v0.22.0 // indirect
	github.com/nats-io/nats.go v1.11.0
	github.com/nats-io/stan.go v0.9.0
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/oschwald/maxminddb-golang v1.8.0
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.22.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/sirupsen/logrus v1.8.1
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/smallnest/gofsm v0.0.0-20190306032117-f5ba1bddca7b
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/stretchr/testify v1.7.0
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125
	github.com/vmware/govmomi v0.26.0
	github.com/wangtuanjie/ip17mon v1.5.2
	github.com/xtaci/kcp-go/v5 v5.6.1
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9 // indirect
	github.com/zommage/cron v0.0.0-20180918061821-210507a89644
	go.elastic.co/apm/module/apmgin v1.14.0
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	go.etcd.io/etcd/client/pkg/v3 v3.5.0
	go.etcd.io/etcd/raft/v3 v3.5.0
	go.etcd.io/etcd/server/v3 v3.5.0
	go.mongodb.org/mongo-driver v1.5.3
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210716203947-853a461950ff
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.0 // indirect
	gopkg.in/bufio.v1 v1.0.0-20140618132640-567b2bfa514e // indirect
	gopkg.in/mcuadros/go-syslog.v2 v2.3.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/redis.v2 v2.3.2
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.21.10
	gotest.tools v2.2.0+incompatible
)

// bug from etcd
// refer: https://www.cnblogs.com/anmutu/p/etcd.html
// google.golang.org/grpc 1.26 之後的版本是不支持 clientv3的，所以要把它改成 1.26.0
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
