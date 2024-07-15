package PackageProject

import (
	"net/http"

	"air-drop/cmd/internal/logic/admin/PackageProject"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddOrUpdatePackageProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOrUpdatePackageProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := PackageProject.NewAddOrUpdatePackageProjectLogic(r.Context(), svcCtx)
		resp, err := l.AddOrUpdatePackageProject(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
