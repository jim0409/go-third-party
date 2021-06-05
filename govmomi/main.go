// https://hawksnowlog.blogspot.tw/2017/06/fetch-vapps-with-govmomi.html
package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
)

var (
	envURL         = "https://172.31.17.100/sdk"
	urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", envURL)
	urlFlag        = flag.String("url", envURL, urlDescription)

	envInsecure         = true
	insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", envInsecure)
	insecureFlag        = flag.Bool("insecure", envInsecure, insecureDescription)
)

func main() {
	// vCenter への接続
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	flag.Parse()
	u, err := url.Parse(*urlFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	u.User = url.UserPassword("agent.test", "agent.test")
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c.Client.Client.Version)
	// fmt.Println(c.Client.Client.UserAgent)

	// データセンターの取得
	f := find.NewFinder(c.Client, true)

	dc, err := f.Datacenter(ctx, "DiskProphet")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.SetDatacenter(dc)
	fmt.Println(dc)

	vas, err := f.VirtualMachineList(ctx, "*")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f1, _ := vas[0].QueryConfigTarget(ctx)
	for _, t := range f1.Network {
		fmt.Printf("%+v", t)
	}
}
