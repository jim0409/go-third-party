package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
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

	// TODO: refactor
	sess, err = mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}

	sess.Login(&mgo.Credential{
		Username: db.username,
		Password: db.password,
	})

	return client, nil
}

type OPDB struct {
	Client *mongo.Client
}

type method interface {
	Create(string, int) error // account, money
}

func (db *OPDB) Create(account string, money int) error {

	// var db *mongo.Database

	// Create a "users" collection with a JSON schema validator. The validator will ensure that each document in the
	// collection has "name" and "age" fields.
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "age"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "the name of the user, which is required and must be a string",
			},
			"age": bson.M{
				"bsonType":    "int",
				"minimum":     money,
				"description": "the age of the user, which is required and must be an integer >= 18",
			},
		},
	}
	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}
	opts := options.CreateCollection().SetValidator(validator)

	return db.Client.Database("mongo").CreateCollection(context.TODO(), account, opts)

}

// refer: https://stackoverflow.com/a/58413538
// refer: https://stackoverflow.com/a/30342840

var sess *mgo.Session
var client *mongo.Client

type Data struct {
	// ID      primitive.ObjectID `bson:"_id"`
	ID      string `bson:"_id"`
	Counter int
}

func (db *OPDB) Bulk(account string, money int) error {

	if err := sess.Login(&mgo.Credential{Username: "root", Password: "password"}); err != nil {
		return err
	}

	c := sess.DB("mongo").C("bulk")
	if _, err := c.RemoveAll(bson.M{}); err != nil {
		return err
	}

	bulk := c.Bulk()

	bulk.Insert(&Data{ID: account, Counter: money})

	_, err := bulk.Run()
	return err
}
