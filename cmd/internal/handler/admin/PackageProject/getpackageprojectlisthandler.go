package PackageProject

import (
	"air-drop/pkg/result"
	"net/http"

	"air-drop/cmd/internal/logic/admin/PackageProject"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPackageProjectListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminGetPackageProjectListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpErrorResult(r.Context(), w, err)
			return
		}

		l := PackageProject.NewGetPackageProjectListLogic(r.Context(), svcCtx)
		resp, err := l.GetPackageProjectList(&req)
		if err != nil {
			result.HttpErrorResult(r.Context(), w, err)
		} else {
			result.HttpSuccessResult(r.Context(), w, resp)
		}
	}
}
