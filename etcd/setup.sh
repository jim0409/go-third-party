#!/bin/bash

# dig k8s ip and cat and config.go file

ip1=`dig etcd1.dev-morse.svc.cluster.local +short`
ip2=`dig etcd2.dev-morse.svc.cluster.local +short`
ip3=`dig etcd3.dev-morse.svc.cluster.local +short`


echo "$ip1:2380,$ip2:2380,$ip3:2380"
