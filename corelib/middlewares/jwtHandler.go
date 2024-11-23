package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/vivekprm/go-microservices-full-stack/corelib/models"
)

type JwtHandler struct {
	Next http.Handler
}

type Exception struct {
	Message string
}

func (jh *JwtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var header = r.Header.Get("x-access-token") //Grab the token from the header
	header = strings.TrimSpace(header)

	if header == "" {
		//Token is missing, returns with error code 403 Unauthorized
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
		return
	}
	tk := &models.Token{}

	_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return
	}
	ctx := context.WithValue(r.Context(), "user", tk)
	jh.Next.ServeHTTP(w, r.WithContext(ctx))
}
