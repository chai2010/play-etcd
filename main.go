package main

import (
	_ "github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang/protobuf/proto"
	_ "google.golang.org/grpc"
	_ "gopkg.in/gorp.v2"
)

func main() {
	println("hi etcd v3")
}
