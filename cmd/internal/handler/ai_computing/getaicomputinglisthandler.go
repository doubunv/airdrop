package ai_computing

import (
	"air-drop/pkg/result"
	"net/http"

	"air-drop/cmd/internal/logic/ai_computing"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAiComputingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAiComputingListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := ai_computing.NewGetAiComputingListLogic(r.Context(), svcCtx)
		resp, err := l.GetAiComputingList(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
