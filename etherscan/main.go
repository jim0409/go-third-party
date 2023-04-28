package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Account string   `json:"account"`
	Targets []string `json:"targets"`
}

func main() {
	route := gin.New()

	corscfg := cors.DefaultConfig()
	corscfg.AllowAllOrigins = true
	route.Use(cors.New(corscfg))
	route.POST("etherscan", TxVerifiedHandler)

	httpSrv := &http.Server{
		Addr:    ":8000",
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("http listen : %v\n", err)
			panic(err)
		}
	}()

	select {}
}

func TxVerifiedHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		ErResp(c, http.StatusBadRequest, ReadBodyErr, err)
		return
	}

	reqbody := &RequestBody{}
	err = json.Unmarshal(body, reqbody)
	if err != nil {
		ErResp(c, http.StatusBadRequest, InvalidJson, err)
		return
	}

	// account := c.Param("account")
	account := reqbody.Account
	if account == "" {
		ErResp(c, http.StatusBadRequest, LackAccountInfo, nil)
		return
	}

	// target := c.Query("target")
	targets := reqbody.Targets
	if targets == nil {
		ErResp(c, http.StatusBadRequest, LackTargetInfo, nil)
		return
	}

	// TODO: consider take cache ?
	itxlist, err := GetAccountTxList(account)
	if err != nil {
		ErResp(c, http.StatusInternalServerError, "", err)
		return
	}

	// TODO: consider cache update ?
	res := itxlist.SearchTargetTx(targets)

	OkResp(c, http.StatusOK, MatchAcccountTx, gin.H{"account": account, "tx": res})
}
