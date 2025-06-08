package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/internal/svc"
	"github.com/zyyujkkj/shorturl-service/rpc/transform/transform"
	"github.com/zyyujkkj/shorturl-service/rpc/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortUrlLogic {
	return &GetShortUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShortUrlLogic) GetShortUrl(in *transform.GetShortUrlRequest) (*transform.GetShortUrlResponse, error) {
	hash := md5.Sum([]byte(in.Url))
    key := hex.EncodeToString(hash[:])[:6]  // 取前6位做短码
    _, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
        ShortUrl: key,
        Url:      in.Url,
    })
    if err != nil {
        return nil, err
    }
    return &transform.GetShortUrlResponse{
        ShortUrl: key,
    }, nil
}
