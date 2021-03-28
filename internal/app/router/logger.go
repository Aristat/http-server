package router

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5/middleware"
)

func loggerMiddleware(l *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.WithFields(logrus.Fields{
					"proto":  r.Proto,
					"path":   r.URL.Path,
					"lat":    time.Since(t1),
					"status": ww.Status(),
					"size":   ww.BytesWritten(),
					"reqId":  middleware.GetReqID(r.Context()),
				}).Info()
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
