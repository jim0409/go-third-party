#!/bin/bash

file=$1
order=$2
total=$3
md5_value=`md5 $file | awk '{print $4}'`

curl -F "myFile=@$file" -H "username: jim" "http://127.0.0.1:8000/file/upload?md5value=$md5_value&filename=$file&chunkorder=$order&totalchunks=$total"

