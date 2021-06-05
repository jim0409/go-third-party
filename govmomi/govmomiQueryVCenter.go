package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type Neo4j struct {
	Urls               string
	InsecureSkipVerify bool
}

type nodeInfo struct {
	NodeNum  string
	DomainID string
	Name     string
	Labels   interface{}
	Types    string
}

func vCenterVmName(neo4j Neo4j) map[int]nodeInfo {

	// 製作一個ctx當作紀錄點
	ctx, _ := context.WithCancel(context.Background())

	// 給定要查詢的網址以及對應的使用者名稱及密碼
	u, _ := url.Parse(neo4j.Urls)
	u.User = url.UserPassword("agent.test", "agent.test")

	// 建立govmomi的新使用者 govmomi.NewClient
	c, _ := govmomi.NewClient(ctx, u, neo4j.InsecureSkipVerify)

	// 建立一個view.NewManager後面可以利用它來查詢nodes
	viewNewManager := view.NewManager(c.Client)
	ContainView, _ := viewNewManager.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	defer ContainView.Destroy(ctx)

	// 使用mo.HostSystem將ctx路徑下的HostSystem的summary紀錄給hss
	var hss []mo.HostSystem
	_ = ContainView.Retrieve(ctx, []string{"HostSystem"}, []string{"summary", "datastore", "vm", "configManager", "systemResources"}, &hss)

	// 打印出來確認hosts名稱
	for _, hs := range hss {
		fmt.Printf("%s\n", hs.Summary.Config.Name)
		fmt.Printf("%v\n", hs.Datastore)

		// fmt.Println(hs.SystemResources)
	}

	fmt.Println("------------above is host IP---------------")

	// ====================================================================================================== //
	// 建立一個Finder
	f := find.NewFinder(c.Client, true)

	// 使用DatacenterList去找尋指定路徑"*"下的datacenterList
	datacenterList, _ := f.DatacenterList(ctx, "*")

	for i := 0; i < len(datacenterList); i++ {
		fmt.Println(datacenterList[i].ObjectName(ctx))
	}
	fmt.Println("----------above would list vmware VMDataCenter-----------")

	objectNameOfDatacenter, _ := datacenterList[1].ObjectName(ctx)

	dc, _ := f.Datacenter(ctx, objectNameOfDatacenter)

	// 設定finder的datacenter為dc
	f.SetDatacenter(dc)

	vas, _ := f.VirtualMachineList(ctx, "*")

	hostsTest, _ := f.HostSystemList(ctx, "*")
	// fmt.Println(hostsTest) //[HostSystem:host-223 @ /DiskProphet/host/172.31.17.88/172.31.17.88 HostSystem:host-113 @ /DiskProphet/host/vSAN/172.31.17.92 HostSystem:host-234 @ /DiskProphet/host/vSAN/172.31.17.96 HostSystem:host-54 @ /DiskProphet/host/vSAN/172.31.17.94]
	fmt.Println(hostsTest[3].ObjectName(ctx))
	fmt.Println("================")

	dsTest, _ := hostsTest[3].ConfigManager().StorageSystem(ctx)
	// fmt.Println(dsTest)

	var hssTT mo.HostStorageSystem

	_ = dsTest.Properties(ctx, dsTest.Reference(), nil, &hssTT)

	for _, e := range hssTT.StorageDeviceInfo.ScsiLun {
		fmt.Println(e.GetScsiLun().DeviceName)
	}

	fmt.Println("==================above is the host's vmdisk==================")

	// 打印確認該datacenter下有幾個datastores
	i, _ := f.DatastoreList(ctx, "*")
	for index := 0; index < len(i); index++ {
		objectNameOfDatastores, _ := i[index].ObjectName(ctx)
		fmt.Printf("%s\n", objectNameOfDatastores)
	}

	fmt.Println("----------above would list vmware DataCenter's datastores-----------")

	s := make(map[int]nodeInfo, len(vas))

	for index, va := range vas {
		var o mo.VirtualMachine
		_ = vas[index].Properties(ctx, vas[index].Reference(), []string{"snapshot"}, &o)
		if o.Snapshot != nil {
			fmt.Println("index:", index, " va.Name:", va.Name())
			fmt.Println("check leaf")
			check(o.Snapshot.RootSnapshotList)
		}
		keyString := fmt.Sprintf("n%d", index)
		if index == 0 {
			s[index] = nodeInfo{
				NodeNum:  keyString,
				DomainID: va.Name(),
				Name:     va.Name(),
				Types:    "nodes",
				Labels:   "nodes",
			}
		} else {
			continue
		}
	}
	return s
}

func main() {
	neo4jTest := Neo4j{
		Urls:               "https://172.31.17.100/sdk",
		InsecureSkipVerify: true,
	}
	fmt.Println(vCenterVmName(neo4jTest))
}

func check(leaf []types.VirtualMachineSnapshotTree) string {
	if leaf[0].ChildSnapshotList != nil {
		for index := range leaf[0].ChildSnapshotList {
			fmt.Println("the leaf is", index)
			check(leaf[0].ChildSnapshotList)
		}
	} else {
		fmt.Println("no other leaf left")
	}
	return "nothing"
}
