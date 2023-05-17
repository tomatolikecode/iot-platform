package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf      rest.RestConf
	RpcClientConf zrpc.RpcClientConf
}
