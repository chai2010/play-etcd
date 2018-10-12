package main

import (
	_ "github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/yuin/gopher-lua"
	_ "golang.org/x/crypto/bcrypt"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/net/proxy"
	_ "golang.org/x/text"
	_ "google.golang.org/grpc"
	_ "gopkg.in/gorp.v2"
)

func main() {
	println("hi etcd v3")
}
