package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	RpcServerConf zrpc.RpcServerConf
	RpcClientConf zrpc.RpcClientConf
	Mqtt          struct {
		Broker   string
		ClientID string
		Password string
	}
}
