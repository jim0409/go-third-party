package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("tags")
	collection := quickstartDatabase.Collection("user")

	// filter := bson.M{"lineUserID": "U52cb972d698a42338f63439179206c4e"}
	/*
		- https://docs.mongodb.com/manual/reference/operator/query/
		善加利用mongo的語法可以優化工作效率
	*/
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	// cursor, err := collection.Find(ctx, bson.M{"lineUserID": "U52cb972d698a42338f63439179206c4e"})
	if err != nil {
		log.Fatal(err)
	}

	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(episodes)  // return []map[interface]interface{}

	f, _ := os.Create("./tags.log")

	for i, j := range episodes {
		// if j["tags"] != nil { // 資料型別為private.M，所以不能用非nil來算
		if fmt.Sprintf("%v", j["tags"]) != "[]" {
			f.WriteString(fmt.Sprintf("%d__%v\n", i, j)) // 將資料寫到檔案內
			// fmt.Println(i, j)

			// objTags := j["tags"]
			// fmt.Println(objTags)
			// fmt.Println(reflect.TypeOf(objTags))
			// fmt.Println(fmt.Sprintf("%v", objTags) == "[]")
		}
	}

}
