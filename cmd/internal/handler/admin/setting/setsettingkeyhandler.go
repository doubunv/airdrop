package setting

import (
	"net/http"

	"air-drop/cmd/internal/logic/admin/setting"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetSettingKeyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetSettingKeyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := setting.NewSetSettingKeyLogic(r.Context(), svcCtx)
		resp, err := l.SetSettingKey(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
