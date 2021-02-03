package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// see https://blog.questionable.services/article/guide-logging-middleware-go/
// for the fuckery that follows.

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

// Log wraps original request with logs
func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := wrapResponseWriter(w)
		start := time.Now()
		h.ServeHTTP(wrapped, r)
		log.Info().
			Int("status", wrapped.status).
			Dur("elapsed_ms", time.Now().Sub(start)).
			Msg(fmt.Sprintf("%v", r.Method))
	})
}
