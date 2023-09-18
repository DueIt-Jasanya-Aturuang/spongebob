package middleware

import (
	"net/http"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/config"
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
