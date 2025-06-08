package logic

import (
	"context"
	"github.com/zyyujkkj/shorturl-service/api/internal/svc"
	"github.com/zyyujkkj/shorturl-service/api/internal/types"	
	"github.com/zeromicro/go-zero/core/logx"
	//"github.com/zyyujkkj/shorturl-service/rpc/transform/transformclient"
)

import transform "github.com/zyyujkkj/shorturl-service/rpc/transform/transformclient"


type GetShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortUrlLogic {
	return &GetShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShortUrlLogic) GetShortUrl(req *types.GetShortUrlRequest) (resp *types.GetShortUrlResponse, err error) {
	ret, err := l.svcCtx.Transformer.GetShortUrl(l.ctx, &transform.GetShortUrlRequest{
		Url: req.Url,
	})
    if err != nil {
        return nil, err
    }
    return &types.GetShortUrlResponse{
        ShortUrl: ret.ShortUrl,
    }, nil
}	
