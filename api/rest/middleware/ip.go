package middleware

import (
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/config"
)

func IPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr != "27.112.78.47" && config.AppStatus == "production" {
			helper.ErrorResponseEncode(w, _error.HttpErrString(string(response.CM05), response.CM05))
			return
		}
		next.ServeHTTP(w, r)
	})
}
