#!/usr/bin/env bash

web="http://www.ziyuandaigou7.com/page/1?v=0.12322"
times="10"

dir=$(cd $(dirname $0); pwd)
mainFile=$(cd ${dir}/../; pwd)/main.go

if [ -f $mainFile ]; then
  go run $mainFile $web $times
else
  echo $mainFile"文件不存在"
fi
