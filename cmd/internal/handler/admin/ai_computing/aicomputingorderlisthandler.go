package ai_computing

import (
	"air-drop/pkg/result"
	"net/http"

	"air-drop/cmd/internal/logic/admin/ai_computing"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AiComputingOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminAiComputingOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := ai_computing.NewAiComputingOrderListLogic(r.Context(), svcCtx)
		resp, err := l.AiComputingOrderList(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
