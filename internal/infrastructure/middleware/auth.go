package middleware

import (
	"context"
	"net/http"
)

const CurrentData = "currentData"

type UserCtx struct {
	UserId string `json:"token"`
}

type DataCtx struct {
	UserC UserCtx
}

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeaderValue := r.Header.Get("token")
		if tokenHeaderValue == "" {
			next.ServeHTTP(w, r)
			return
		}
		userCtx := UserCtx{UserId: tokenHeaderValue}
		dataCtx := DataCtx{UserC: userCtx}
		ctx := context.WithValue(r.Context(), CurrentData, dataCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
