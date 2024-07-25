package person

import (
	"air-drop/pkg/result"
	"net/http"

	"air-drop/cmd/internal/logic/person"
	"air-drop/cmd/internal/svc"
)

func NonceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := person.NewNonceLogic(r.Context(), svcCtx)
		resp, err := l.Nonce()
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
