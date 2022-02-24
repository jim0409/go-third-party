package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/arangodb/go-driver"
	"github.com/stretchr/testify/assert"
)

func TestQueryInsert(t *testing.T) {
	db, err := NewArangoDB("http://127.0.0.1:8529")
	assert.Nil(t, err)

	err = QueryInsertCollection("test", db)
	assert.Nil(t, err)

	err = QueryInsertCollection("test2", db)
	assert.Nil(t, err)

	fmt.Println("success")
}

func QueryInsertCollection(col string, db driver.Database) error {
	ctx := context.TODO()
	if found, err := db.CollectionExists(ctx, col); err != nil || found {
		if err != nil {
			fmt.Println(err)
		}
		// return nil if founded
		return err
	}

	if _, err := db.CreateCollection(ctx, col, nil); err != nil {
		return err
	}

	return nil
}

func QueryCount(db driver.Database, col string) (int64, error) {
	ctx := context.Background()
	collection, err := db.Collection(ctx, col)
	if err != nil {
		return 0, err
	}
	count, err := collection.Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
func TestQueryCount(t *testing.T) {
	db, err := NewArangoDB("http://127.0.0.1:8529")
	assert.Nil(t, err)
	count, err := QueryCount(db, "AdminAccounts")
	assert.Nil(t, err)
	fmt.Println(count)
}

// https://www.arangodb.com/docs/stable/drivers/go-example-requests.html#querying-documents-one-document-at-a-time
// QueryOrm : arangodb 查詢使用 ORM，需要使用指定的 context，或者透過 driver.IsNoMoreDocuments 來撈取資料
func QueryOrm(db driver.Database, rawQueryStr string, bindVars map[string]interface{}) (*[]interface{}, error) {
	// ctx := driver.WithQueryCount(context.Background())
	ctx := context.Background()
	cursor, err := db.Query(ctx, rawQueryStr, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	fmt.Printf("query yields %d\n", cursor.Count())
	for {
		var doc interface{}
		meta, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		fmt.Println("---------", meta)
		fmt.Println("=========", doc)
	}

	return nil, nil
}

func TestQueryOrm(t *testing.T) {
	db, err := NewArangoDB("http://127.0.0.1:8529")
	assert.Nil(t, err)

	qstr := `FOR d IN AdminAccounts RETURN d`
	data, err := QueryOrm(db, qstr, nil)
	assert.Nil(t, err)

	fmt.Println(data)
}
