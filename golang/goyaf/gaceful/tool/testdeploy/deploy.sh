#!/bin/bash


basepath=$(cd `dirname $0`; pwd)
programDir=${basepath}"/../../"

$programDir/tool/compile.sh

#$programDir/tool/testdeploy/rm.sh

cd $programDir/bin

tar -jcf gaceful-linux64.tar.bz gaceful-linux64

scp $programDir/bin/gaceful-linux64.tar.bz root@172.16.27.135:/root/go

#$programDir/tool/testdeploy/reload.sh