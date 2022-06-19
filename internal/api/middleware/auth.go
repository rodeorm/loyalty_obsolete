package middleware

import (
	"net/http"

	cookie "loyalty/internal/api/cookie"
)

// AuthMiddleware выполняется для проверки аутентифицирован ли пользователь
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userKey, err := cookie.GetUserKeyFromCoockie(r)
		if err != nil {
			http.Redirect(w, r, "/forbidden", http.StatusFound)
			return
		}
		if userKey == "" {
			http.Redirect(w, r, "/forbidden", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
