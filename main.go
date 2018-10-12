package main

import (
	_ "github.com/coreos/etcd/clientv3"
	_ "github.com/golang/protobuf/proto"
	_ "google.golang.org/grpc"
)

func main() {
	println("hi etcd v3")
}
