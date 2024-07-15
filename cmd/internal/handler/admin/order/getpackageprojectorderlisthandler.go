package order

import (
	"net/http"

	"air-drop/cmd/internal/logic/admin/order"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPackageProjectOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPackageProjectOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewGetPackageProjectOrderListLogic(r.Context(), svcCtx)
		resp, err := l.GetPackageProjectOrderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
