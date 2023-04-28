package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testAccount  = "0x755229B87cEc1d55c8b1Aa0c66C1be4D8136E8CF"
	targetTxHash = "0x5f432212023ecd0019aca196d561e9ac96a95dee5f38ddfeae89647df45d6126"
)

func TestCheckAccountTxList(t *testing.T) {
	itxlist, err := GetAccountTxList(testAccount)
	assert.Nil(t, err)

	fmt.Println(itxlist.SearchTargetTx([]string{targetTxHash}))
}
