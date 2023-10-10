package middleware

import (
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/parse"
)

func IPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr != "27.112.78.47" && infra.AppStatus == "production" {
			parse.ErrorResponseEncode(w, _error.HttpErrString(string(response.CM05), response.CM05))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newToken := r.Header.Get("Authorization")
		if newToken != "" {
			w.Header().Set("Authorization", newToken)
		}
		next.ServeHTTP(w, r)
	})
}

func CheckApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey != infra.AppApiKeyAccount {
			log.Warn().Msgf("invalid api key | key : %s", apiKey)
			parse.ErrorResponseEncode(w, _error.HttpErrString("forbidden", response.CM05))
			return
		}
		next.ServeHTTP(w, r)
	})
}
