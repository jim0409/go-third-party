#!/bin/bash
######### WARNNING #########
# this is anti-container!! #
############################

function binary() {
	CGO_ENABLED=0 GOOS=linux go build http_add.go db.go	
}

function build_docker_img() {
	docker build -t reg.paradise-soft.com.tw:5000/xread .
}


function clean() {
	rm http_add
}

function push_docker_img() {
	docker push reg.paradise-soft.com.tw:5000/xread
}

binary
build_docker_img
clean
push_docker_img