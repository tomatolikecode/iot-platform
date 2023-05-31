package main

import (
	"flag"
	"fmt"

	"github.com/iot-platform/device/internal/config"
	"github.com/iot-platform/device/internal/mqtt"
	"github.com/iot-platform/device/internal/server"
	"github.com/iot-platform/device/internal/svc"
	"github.com/iot-platform/device/types/device"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/device.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewDeviceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		device.RegisterDeviceServer(grpcServer, svr)

		if c.RpcServerConf.Mode == service.DevMode || c.RpcServerConf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	go mqtt.NewMqttServer(c.Mqtt.Broker, c.Mqtt.ClientID, c.Mqtt.Password)

	fmt.Printf("Starting rpc server at %s...", c.RpcServerConf.ListenOn)
	s.Start()
}
