#!/bin/bash

protoc -I=. --go_out=. proto/basic.proto util/util.proto
