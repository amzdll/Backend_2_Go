package middleware

import (
	"fmt"
	"github.com/amzdll/backend_2_go/src/internal/app/logger"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Logger struct {
	logger *logger.Logger
}

func NewLogger(logger *logger.Logger) *Logger {
	return &Logger{logger: logger}
}

// need decompose
func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		l.logger.Info(
			fmt.Sprintf("method=%s path=%s remote_addr=%s request_id=%s status=%d",
				r.Method, r.URL.Path, r.RemoteAddr, middleware.GetReqID(r.Context()), ww.Status()),
		)
	})
}
