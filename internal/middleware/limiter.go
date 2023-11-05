package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
)

func Limit(next http.Handler) http.Handler {
	var limiter = rate.NewLimiter(100, 100)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, "429 "+http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
