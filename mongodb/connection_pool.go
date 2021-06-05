package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	MAX_CONNECTION     = 10
	INITIAL_CONNECTION = 4
	AVAILABLE          = false
	USED               = true
)

var (
	mu sync.RWMutex
	cp ClientPool
)

type mongodata struct {
	client *mongo.Client
	pos    int
	flag   bool
}

/*
	clientList: the client pool
	clientAvailable: the availabl flag, means the location and available flag in the client pool
	size: the size of allocated client pool (less than MAX_CONNECTION)
*/
type ClientPool struct {
	clientList [MAX_CONNECTION]mongodata
	// clientList make([]]mongodata, MAX_CONNECTION)
	size int
	DBConfig
}

func disconnectToMongoDB(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}

func (cp *ClientPool) allocateCToPool(pos int) error {
	var err error
	if cp.clientList[pos].client, err = cp.DBConfig.connectToMongoDB(); err != nil {
		return fmt.Errorf("Encounter error while allocate to pool %v", err)
	}

	cp.clientList[pos].flag = USED
	cp.clientList[pos].pos = pos

	return err
}

//apply a connection from the pool
func (cp *ClientPool) getCToPool(pos int) {
	cp.clientList[pos].flag = USED
}

//free a connection back to the pool
func (cp *ClientPool) putCBackPool(pos int) {
	cp.clientList[pos].flag = AVAILABLE
}

//program apply a database connection
func GetClient() (*mongodata, error) {
	mu.RLock()
	for i := 1; i < cp.size; i++ {
		if cp.clientList[i].flag == AVAILABLE {
			return &cp.clientList[i], nil
		}
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()
	if cp.size < MAX_CONNECTION {
		if err := cp.allocateCToPool(cp.size); err != nil {
			return nil, err
		}

		pos := cp.size
		cp.size++
		return &cp.clientList[pos], nil
	} else {
		return nil, fmt.Errorf("DB pooling is fulled")
	}
}

//program release a connection
func ReleaseClient(mongoclient *mongodata) {
	mu.Lock()
	cp.putCBackPool(mongoclient.pos)
	mu.Unlock()
}

func main() {
	// 1. setup db connection settings
	dbconfig := &DBConfig{
		username:              "root",
		password:              "password",
		address:               "127.0.0.1:27017",
		maxConnectionIdleTime: 5,
		maxPoolSize:           200,
	}
	cp.DBConfig = *dbconfig

	// append connections to connections_pool
	var err error
	for size := 0; size < INITIAL_CONNECTION || size < MAX_CONNECTION; size++ {
		// 2. generate client ~
		// 3. append client to connection pools
		// 4. work around
		if err = cp.allocateCToPool(size); err != nil {
			panic(err)
		}
	}

	/*
		TODO:
		rewrite as ... send task to channel and client would handle all the channel task
	*/
	// mgopc mongo pool client
	mgopc, err := GetClient()
	if err != nil {
		panic(err)
	}

	quickDB := mgopc.client.Database("tags")
	collection := quickDB.Collection("user")

	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var res []bson.M
	if err = cursor.All(context.TODO(), &res); err != nil {
		panic(err)
	}

	f, _ := os.Create("./tags.log")
	for i, j := range res {
		// fmt.Println(i, j)
		if fmt.Sprintf("%v", j["tags"]) != "[]" {
			f.WriteString(fmt.Sprintf("%d__%v\n", i, j)) // 將資料寫到檔案內
		}
	}
}

type DBConfig struct {
	username              string
	password              string
	address               string // where address would be "ip1:port1,ip2:port2,ip3:port3"
	maxConnectionIdleTime time.Duration
	maxPoolSize           uint64
}

/*
	connect to mongoDB address with specific username and password
	where address would be "ip1,ip2,ip3" as cluster annotation
*/
func (db *DBConfig) connectToMongoDB() (*mongo.Client, error) {
	opts := options.Client()
	// any of username or password empty would by pass auth
	if db.username != "" || db.password != "" {
		opts.Auth = &options.Credential{
			Username: db.username,
			Password: db.password,
		}
	}

	// TODO: a better formater checker for address
	if db.address == "" {
		return nil, fmt.Errorf("Invalid address with value should be 192.168.xxx.xxx:21017 instead of %v", db.address)
	}

	opts.ApplyURI("mongodb://" + db.address)

	// db macConnectionIdleTime must greater than 0
	if db.maxConnectionIdleTime <= 0 {
		db.maxConnectionIdleTime = 5 // a default value 5 would be given if it is illegal
	}
	opts.SetMaxConnIdleTime(db.maxConnectionIdleTime)

	// db macPoolSize must greater than 0
	if db.maxPoolSize <= 0 {
		db.maxPoolSize = 200 // a default value 200 would be given if it is illegal
	}
	opts.SetMaxPoolSize(db.maxPoolSize)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo server with err : %v", err)
	}

	// cancel is a self defined default timeout : 20 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect mongod ctx failed : %v", err)
	}

	return client, nil
}
