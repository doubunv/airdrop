package middleware

import (
	"net/http"
)

func ActiveUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//userAddress := utils.GetTokenAddress(r.Context())
		//if userAddress != "" {
		//	err := service.NewUserStatisticsService(r.Context(), svc.SCtx).AddActiveUser(userAddress, time.Now())
		//	if err != nil {
		//		logx.Errorf("ActiveUserMiddleware: %s", err.Error())
		//	}
		//}

		next(w, r)
	}
}
