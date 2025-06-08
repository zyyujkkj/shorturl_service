package main

import (
	"flag"
	"fmt"

	"github.com/zyyujkkj/shorturl-service/rpc/transform/internal/config"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/internal/server"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/internal/svc"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/transform.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		transform.RegisterTransformServer(grpcServer, server.NewTransformServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
