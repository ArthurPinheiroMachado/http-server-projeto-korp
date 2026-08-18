package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *statusResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *statusResponseWriter) Write(body []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	return rw.ResponseWriter.Write(body)
}

func MetricsMiddleware(counter *prometheus.CounterVec, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := &statusResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next(rw, r)

		path := r.URL.Path
		if route := mux.CurrentRoute(r); route != nil {
			if tpl, err := route.GetPathTemplate(); err == nil {
				path = tpl
			}
		}

		counter.WithLabelValues(
			r.Method,
			path,
			strconv.Itoa(rw.statusCode),
		).Inc()

	}
}
