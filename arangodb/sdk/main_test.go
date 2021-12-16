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
