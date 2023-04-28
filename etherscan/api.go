package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const gateway = "https://api.etherscan.io/api"

// according to account address get related txlist
func GetAccountTxList(account string) (ITxList, error) {
	action := "txlist"
	module := "account"
	resp, err := http.Get(fmt.Sprintf("%v?module=%v&action=%v&address=%v", gateway, module, action, account))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	txlist := &AccountTxList{}
	err = json.Unmarshal(bs, &txlist)
	if err != nil {
		return nil, err
	}

	return txlist, nil
}

type ITxList interface {
	SearchTargetTx(targets []string) []string
}

func (a *AccountTxList) SearchTargetTx(targets []string) []string {
	txs := []string{}

	targetmaps := make(map[string]bool)
	for _, target := range targets {
		targetmaps[target] = true
	}

	for _, res := range a.Result {
		if _, ok := targetmaps[res.Hash]; ok {
			if res.IsError == "0" {
				txs = append(txs, res.Hash)
			}
		}
	}

	return txs
}

type AccountTxList struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// TODO: refactor with map?
	// Result []map[string]interface{} `result`
	Result []struct {
		BlockNumber       string `json:"blockNumber"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		TransactionIndex  string `json:"transactionIndex"`
		From              string `json:"from"`
		To                string `json:"to"`
		Value             string `json:"value"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		IsError           string `json:"isError"`
		TxreceiptStatus   string `json:"txreceipt_status"`
		Input             string `json:"input"`
		ContractAddress   string `json:"contractAddress"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		GasUsed           string `json:"gasUsed"`
		Confirmations     string `json:"confirmations"`
		MethodID          string `json:"methodId"`
		FunctionName      string `json:"functionName"`
	} `json:"result"`
}
