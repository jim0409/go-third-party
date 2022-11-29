package main

import (
	"fmt"
	"testing"
)

func TestDomains(t *testing.T) {

	// TODO: assume mock httpServer ?
	m, err := domains("280076842", "e5N8Yzzn7hBo_Fxe7rVEZwvyTfMW4ztQmJG", "PDnDKUFDZjefBUUyTeWGkb")
	assert.NilError(t, err)
	for i, j := range m {
		fmt.Println(i, j)
	}
}
