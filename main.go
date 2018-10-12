package main

import (
	_ "github.com/coreos/etcd/clientv3"
	_ "github.com/dgraph-io/badger"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/koding/multiconfig"
	_ "github.com/pkg/errors"
	_ "github.com/spf13/pflag"
	_ "github.com/yuin/gopher-lua"
	_ "golang.org/x/crypto/bcrypt"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/net/proxy"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/text"
	_ "golang.org/x/tools/godoc/vfs"
	_ "google.golang.org/grpc"
	_ "gopkg.in/gorp.v2"
)

func main() {
	println("hi etcd v3")
}
