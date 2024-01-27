package main

import (
	"fmt"
	"rpc-template/internal/rabbitmq/consumer"
	server "rpc-template/internal/server/userdetailservice"
	"sync"

	"rpc-template/internal/config"
	"rpc-template/internal/svc"
	"rpc-template/pb/rpcTemplate"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Init()
	ctx := svc.NewServiceContext(*c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpcTemplate.RegisterUserDetailServiceServer(grpcServer, server.NewUserDetailServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode || c.Mode == service.ProMode {

			wg := new(sync.WaitGroup)
			go func() {
				wg.Add(1)
				consumer.ConsumeDetailListCount(ctx, wg)
			}()

			reflection.Register(grpcServer)
		}

	})

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}