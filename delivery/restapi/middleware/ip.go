package middleware

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"net/http"
)

func IPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr != "27.112.78.47" && config.AppStatus == "production" {
			response.NewError(w, r, model.ErrForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
