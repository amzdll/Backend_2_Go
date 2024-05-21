package middleware

import (
	"fmt"

	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type log interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string, err error)
	Warn(msg string)
	Fatal(msg string, err error)
}

type Logger struct {
	log log
}

func NewLogger(l log) *Logger {
	return &Logger{log: l}
}

func (l *Logger) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		l.log.Info(
			fmt.Sprintf("method=%s path=%s remote_addr=%s request_id=%s status=%d",
				r.Method, r.URL.Path, r.RemoteAddr, middleware.GetReqID(r.Context()), ww.Status()),
		)
	})
}
