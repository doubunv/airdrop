package person

import (
	"air-drop/pkg/result"
	"net/http"

	"air-drop/cmd/internal/logic/person"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InviteCodeChangeSelfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InviteCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := person.NewInviteCodeChangeSelfLogic(r.Context(), svcCtx)
		err := l.InviteCodeChangeSelf(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, nil)
		}
	}
}
