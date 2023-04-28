module go-third-party

go 1.20

// bug from etcd
// refer: https://www.cnblogs.com/anmutu/p/etcd.html
// google.golang.org/grpc 1.26 之後的版本是不支持 clientv3的，所以要把它改成 1.26.0
// replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/Jeffail/gabs v1.4.0
	github.com/LdDl/fiber-long-poll/v2 v2.2.4
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/Shopify/sarama v1.38.1
	github.com/adamwasila/arangodb-adapter v0.0.0-20230310000008-a9ec99915b83
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/appleboy/graceful v0.0.4
	github.com/arangodb/go-driver v1.5.2
	github.com/astaxie/beego v1.12.3
	github.com/aws/aws-lambda-go v1.39.1
	github.com/bluele/slack v0.0.0-20180528010058-b4b4d354a079
	github.com/casbin/casbin/v2 v2.65.2
	github.com/casbin/gorm-adapter/v3 v3.15.0
	github.com/coreos/etcd v3.3.27+incompatible
	github.com/fatih/structs v1.1.0
	github.com/garyburd/redigo v1.6.4
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.9.0
	github.com/go-ini/ini v1.67.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gocolly/colly/v2 v2.1.0
	github.com/gofiber/fiber/v2 v2.42.0
	github.com/gogf/gf v1.16.9
	github.com/gogf/gf/v2 v2.3.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/gopacket v1.1.19
	github.com/google/uuid v1.3.0
	github.com/googollee/go-socket.io v1.7.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/graphql-go/graphql v0.8.0
	github.com/influxdata/influxdb v1.11.0
	github.com/jcuga/golongpoll v1.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo v3.3.10+incompatible
	github.com/lonng/nano v0.5.0
	github.com/looplab/fsm v1.0.1
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/minio/minio-go/v7 v7.0.50
	github.com/mitchellh/mapstructure v1.5.0
	github.com/montanaflynn/stats v0.7.0
	github.com/naoina/toml v0.1.1
	github.com/nats-io/nats.go v1.24.0
	github.com/nats-io/stan.go v0.10.4
	github.com/nsqio/go-nsq v1.1.0
	github.com/olivere/elastic v6.2.37+incompatible
	github.com/oschwald/maxminddb-golang v1.10.0
	github.com/prometheus/client_golang v1.14.0
	github.com/robfig/cron v1.2.0
	github.com/rs/xid v1.4.0
	github.com/rs/zerolog v1.29.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.9.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/smallnest/gofsm v0.0.0-20190306032117-f5ba1bddca7b
	github.com/smartystreets/goconvey v1.7.2
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.2
	github.com/tealeg/xlsx/v3 v3.2.4
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/urfave/cli v1.22.12
	github.com/valyala/fasthttp v1.45.0
	github.com/wangtuanjie/ip17mon v1.5.2
	github.com/xtaci/kcp-go/v5 v5.6.2
	go.elastic.co/apm/module/apmfiber v1.15.0
	go.elastic.co/apm/module/apmgin v1.15.0
	go.etcd.io/etcd/client/pkg/v3 v3.5.7
	go.etcd.io/etcd/raft/v3 v3.5.7
	go.etcd.io/etcd/server/v3 v3.5.7
	go.mongodb.org/mongo-driver v1.11.3
	go.uber.org/fx v1.19.2
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.7.0
	golang.org/x/net v0.8.0
	golang.org/x/time v0.3.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/mcuadros/go-syslog.v2 v2.3.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/redis.v2 v2.3.2
	gorm.io/driver/mysql v1.4.7
	gorm.io/driver/postgres v1.5.0
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.1.8 // indirect
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/clbanning/mxj v1.8.5-0.20200714211355-ff02cfb8ea28 // indirect
	github.com/clbanning/mxj/v2 v2.5.5 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230111030713-bf00bc1b83b6 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/go-licenser v0.3.1 // indirect
	github.com/elastic/go-sysinfo v1.1.1 // indirect
	github.com/elastic/go-windows v1.0.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/frankban/quicktest v1.14.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.19.1 // indirect
	github.com/glebarez/sqlite v1.5.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/grokify/html-strip-tags-go v0.0.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jcchavezs/porto v0.1.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.3 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/klauspost/reedsolomon v1.10.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/nats-io/nats-server/v2 v2.9.15 // indirect
	github.com/nats-io/nats-streaming-server v0.25.3 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/santhosh-tekuri/jsonschema v1.2.4 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20220530130905-52f3993e8d6d // indirect
	github.com/shabbyrobe/xmlwriter v0.0.0-20200208144257-9fca06d00ffa // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/assertions v1.13.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/templexxx/cpu v0.0.9 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/yuin/gopher-lua v0.0.0-20171031051903-609c9cd26973 // indirect
	go.elastic.co/apm v1.15.0 // indirect
	go.elastic.co/apm/module/apmfasthttp v1.15.0 // indirect
	go.elastic.co/apm/module/apmhttp v1.15.0 // indirect
	go.elastic.co/fastjson v1.1.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/pkg/v3 v3.5.7 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/bufio.v1 v1.0.0-20140618132640-567b2bfa514e // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlserver v1.4.1 // indirect
	gorm.io/plugin/dbresolver v1.3.0 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	modernc.org/libc v1.19.0 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.4.0 // indirect
	modernc.org/sqlite v1.19.1 // indirect
)
