package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{ResponseWriter: w}
		// * request info
		method := r.Method
		req_uri := r.RequestURI
		start := time.Now()
		formattedTime := start.Format("Mon, 02 Jan 2006 15:04:05")

		// * Call the next handler
		next.ServeHTTP(wrappedWriter, r)

		// * response info
		statusCode := wrappedWriter.statusCode
		execution_time := time.Since(start)

		fmt.Printf("%s  |%s|  %s  ~  [status/%v]  (%v) \n", formattedTime, method, req_uri, statusCode, execution_time)
	})
}
