package linkproject

import (
	"air-drop/pkg/result"
	"net/http"

	"air-drop/cmd/internal/logic/admin/linkproject"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddOrUpdateLinkProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOrUpdateLinkProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := linkproject.NewAddOrUpdateLinkProjectLogic(r.Context(), svcCtx)
		resp, err := l.AddOrUpdateLinkProject(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
