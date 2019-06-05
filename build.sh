#!/usr/bin/env bash

MAIN_FOLDER="$1"

#设置GOPATH
CURDIR_BASE=`pwd`
CURDIR_LIB="$CURDIR_BASE/src/lib/v"

OLDGOPATH="$GOPATH"
export GOPATH="${CURDIR_BASE}:${CURDIR_LIB}"

echo ${GOPATH}

gofmt -w src

#编译
go_install_shell="go install ${MAIN_FOLDER}"

eval $go_install_shell

#恢复GOPATH
export GOPATH="$OLDGOPATH"

#将产出移到output目录中
mv_shell="mv bin/${MAIN_FOLDER} dist/${MAIN_FOLDER}/bin"

eval $mv_shell

rm -rf bin

echo 'finished'
