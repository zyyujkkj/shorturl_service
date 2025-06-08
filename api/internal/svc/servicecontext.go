package svc

import (
	"github.com/zyyujkkj/shorturl-service/api/internal/config"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/transformclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Transformer transformclient.Transform
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Transformer: transformclient.NewTransform(zrpc.MustNewClient(c.Transform)),
	}
}
