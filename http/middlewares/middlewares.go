package middlewares

import (
	"context"
	"errors"
	"julo-case-study/http/enum"
	"julo-case-study/http/formatter"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func SetContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func CommonAuthJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, status, err := extractToken(r)
		if err != nil {
			formatter.FAIL(w, status, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "ID", claims["ID"])
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func extractToken(r *http.Request) (*jwt.Token, int, error) {
	header := r.Header.Get("Authorization")
	header = strings.TrimSpace(header)

	if header == "" {
		return nil, http.StatusForbidden, errors.New(enum.UnauthorizedUser)
	}

	headerSplit := strings.Split(header, "Token ")
	if len(headerSplit) != 2 {
		return nil, http.StatusForbidden, errors.New(enum.UnauthorizedUser)
	}

	token, err := jwt.Parse(headerSplit[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET")), nil
	})
	if err != nil {
		return nil, http.StatusForbidden, err
	}

	return token, http.StatusOK, err
}
