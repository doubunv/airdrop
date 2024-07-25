package handler

import (
	"air-drop/pkg/result"
	"net/http"
	"os"
	"path/filepath"

	"air-drop/cmd/internal/svc"
)

func swgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := filepath.Abs("cmd/swagger/main.json")

		_, err := os.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				result.HttpErrorResult(r.Context(), w, err)
				return
			}
		}

		http.ServeFile(w, r, s)
	}
}
