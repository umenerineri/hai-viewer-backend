package middleware

import (
	"log"
	"net/http"
	"time"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ここでCORSを設定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")

		// オプションズリクエストに対するプリフライト応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 次のハンドラーを呼び出す
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Received request: %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Handled request: %s %s %s in %v", r.Method, r.RequestURI, r.RemoteAddr, duration)
	})
}
