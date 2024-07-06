package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"air-drop/cmd/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func swgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := filepath.Abs("cmd/swagger/main.json")

		_, err := os.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		}

		http.ServeFile(w, r, s)
	}
}
