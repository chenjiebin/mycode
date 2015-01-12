#!/bin/bash

basepath=$(cd `dirname $0`; pwd)
programDir=${basepath}"/../"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $programDir/bin/linux64 $programDir/src/main.go