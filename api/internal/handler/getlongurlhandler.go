package handler

import (
	"net/http"

	"github.com/zyyujkkj/shorturl-service/api/internal/logic"
	"github.com/zyyujkkj/shorturl-service/api/internal/svc"
	"github.com/zyyujkkj/shorturl-service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetLongUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetLongUrlRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetLongUrlLogic(r.Context(), svcCtx)
		resp, err := l.GetLongUrl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
