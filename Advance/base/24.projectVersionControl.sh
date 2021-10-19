#!/bin/bash

version=""

if [ -f "./VERSION" ]; then
  version=`cat VERSION`
fi

# 判断内容是否为空
if [ -z $version ]; then
  if [ -d ".git" ]; then
    version=`git symbolic-ref HEAD | cut -b 12-`-`git rev-parse HEAD`
  else
    version="unknown"
  fi
fi

# go build project file name
if [ -z $1 ]; then
  echo "project file is required!"
  exit 1
fi
echo "build project file name is ${1}"
go build -ldflags "-X main.Version=$version" -o mainVersionTest $1
