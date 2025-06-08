package logic

import (
	"context"
	//"github.com/zyyujkkj/shorturl-service/rpc/model"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/internal/svc"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLongUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLongUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLongUrlLogic {
	return &GetLongUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLongUrlLogic) GetLongUrl(in *transform.GetLongUrlRequest) (*transform.GetLongUrlResponse, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.Model.FindOneByShortUrl(l.ctx, in.ShortUrl)
	if err != nil {
		return nil, err;
	}
	return &transform.GetLongUrlResponse{
		Url: result.Url,
	}, nil
}
