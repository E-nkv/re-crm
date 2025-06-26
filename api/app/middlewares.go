package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"re-crm/utils"
)

type authCtx struct {
	ID   uint64
	Role string
}

type AuthCtxKeyType string

const AuthCtxKey = AuthCtxKeyType("authCtx")

func (app *App) authMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authJWT")
		if err != nil {
			fmt.Println("👺 no auth cookie found ")
			app.writers.http.Error(w, http.StatusUnauthorized, "no authJWT cookie found")
			return
		}
		token := cookie.Value
		fmt.Println("🚀 the cookie value is: ", cookie.Value)
		claims, err := utils.DecodeJWT(token)
		if err != nil {
			log.Println("invalid jwt: ", err.Error())
			app.writers.http.Error(w, http.StatusUnauthorized, err.Error())
			return
		}
		id, idOk := claims["userID"].(float64)
		role, roleOk := claims["userRole"].(string)
		if !idOk || !roleOk {
			app.writers.http.Error(w, http.StatusBadRequest, "expected id and role")
			return
		}
		if blisted, _ := app.AuthSvc.IsUserBlacklisted(r.Context(), uint64(id)); blisted {
			app.writers.http.Error(w, http.StatusUnauthorized, "user blacklisted")
			return
		}
		ctx := context.WithValue(r.Context(), AuthCtxKey, authCtx{ID: uint64(id), Role: role})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func (app *App) isAdminMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, ok := r.Context().Value(AuthCtxKey).(authCtx)
		if !ok {
			app.writers.http.Error(w, http.StatusBadRequest, "expected authCtx")
			return
		}
		if auth.Role != "admin" {
			app.writers.http.Error(w, http.StatusUnauthorized, "should be admin")
			return
		}
		next.ServeHTTP(w, r)
	})
}
func corsMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:5173" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
