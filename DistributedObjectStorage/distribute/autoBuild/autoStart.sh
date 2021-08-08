#!/bin/bash

export RABBIT_SERVER=amqp://test:test@localhost:5672

LISTEN_ADDRESS=127.0.0.1:8081 STORAGE_ROOT=/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/storage/1 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/dataServer/main.go &
LISTEN_ADDRESS=127.0.0.1:8082 STORAGE_ROOT=/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/storage/2 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/dataServer/main.go &
LISTEN_ADDRESS=127.0.0.1:8083 STORAGE_ROOT=/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/storage/3 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/dataServer/main.go &
LISTEN_ADDRESS=127.0.0.1:8084 STORAGE_ROOT=/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/storage/4 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/dataServer/main.go &
LISTEN_ADDRESS=127.0.0.1:8085 STORAGE_ROOT=/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/storage/5 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/dataServer/main.go &
LISTEN_ADDRESS=127.0.0.1:8086 STORAGE_ROOT=/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/storage/6 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/dataServer/main.go &

LISTEN_ADDRESS=127.0.0.1:8087 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/interfaceServer/main.go &
LISTEN_ADDRESS=127.0.0.1:8088 go run /Users/austsxk/github_sxk_repo/goAdvance/goAdvance/DistributedObjectStorage/distribute/interfaceServer/main.go &