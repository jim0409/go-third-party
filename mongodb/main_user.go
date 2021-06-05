package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

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

var (
	StartMonth = flag.Int("st_mon", 0, "input ur start month")
	StartDay   = flag.Int("st_day", 0, "input ur start day")
	EndMonth   = flag.Int("ed_mon", 0, "input ur end month")
	EndDay     = flag.Int("ed_day", 0, "input ur end day")
)

func main() {
	flag.Parse()
	dbconfig := &DBConfig{
		username:              "root",
		password:              "password",
		address:               "127.0.0.1:27017",
		maxConnectionIdleTime: 5,
		maxPoolSize:           200,
	}

	client, err := dbconfig.connectToMongoDB()
	if err != nil {
		panic(err)
	}

	quickDB := client.Database("tags")
	collection := quickDB.Collection("tag.day")

	var t1 = time.Month(*StartMonth)
	var t2 = time.Month(*EndMonth)

	// fromDate := time.Date(2019, time.November, 1, 0, 0, 0, 0, time.UTC)
	fromDate := time.Date(2019, t1, *EndDay, 0, 0, 0, 0, time.UTC)
	// toDate := time.Date(2020, time.April, 30, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2020, t2, *EndDay, 0, 0, 0, 0, time.UTC)

	filter := bson.M{
		"updateTime": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var res []bson.M
	if err = cursor.All(context.TODO(), &res); err != nil {
		panic(err)
	}

	f, _ := os.Create("./tags.log")
	// f2, _ := os.Create("./tags_u.log")

	var ok bool
	// var lineTags string
	var lineUserID string
	// var userMap map[string]int
	userMap := make(map[string]int)
	// calcualte the whole data
	// count := 0
	for _, j := range res {
		lineUserID, ok = j["lineUserID"].(string)
		if !ok {
			continue
		}
		// lineTags = fmt.Sprintf("%v", j["tag"])

		// 篩選掉tags中為"[]"的人，且限制第一個ID為'U'
		if ok && lineUserID[0] == 'U' {
			convertDate(j)

			// if v, ok := userMap[lineUserID]; ok {
			// if j["compKey"].(int) < v {
			// userMap[lineUserID] = j["compKey"].(int)
			// }
			// } else {
			// userMap[lineUserID] = j["compKey"].(int)
			// }
			userMap[lineUserID] = j["compKey"].(int)

			// f.WriteString(fmt.Sprintf("%d__%v\n", i, j)) // 將資料寫到檔案內
		}

		// if i == 100 {
		// break
		// }

	}

	// fmt.Println(userMap)
	// // log.Println(count)
	dateMap := make(map[int]int)
	for i, j := range userMap {
		dateMap[userMap[i]]++
		f.WriteString(fmt.Sprintf("%v__%v\n", i, j)) // 將資料寫到檔案內

	}

	fmt.Println(dateMap)

}

func convertDate(bM bson.M) {
	timeString := fmt.Sprintf("%v", bM["updateTime"])[:10]
	timeInt, _ := strconv.Atoi(timeString)
	bM["month"] = int(time.Unix(int64(timeInt), 0).Month())
	bM["year"] = int(time.Unix(int64(timeInt), 0).Year())
	bM["compKey"] = int(time.Unix(int64(timeInt), 0).Year())*100 + int(time.Unix(int64(timeInt), 0).Month())
}
