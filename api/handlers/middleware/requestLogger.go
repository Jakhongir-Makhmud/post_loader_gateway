package middleware

import (
	"api-gateway-iman/pkg/logger"
	"net/http"

	"go.uber.org/zap"
)

type Middleware interface {
	LogRequest(http.HandlerFunc) http.HandlerFunc
}

type middleware struct {
	logger logger.Logger
}

func NewMiddleware(l logger.Logger) Middleware {
	return &middleware{
		logger: l,
	}

}

func ApplyMiddleware(h http.HandlerFunc, mw ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {

	for i := range mw {
		h = mw[i](h)
	}
	return h
}

func (m *middleware) LogRequest(req http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		m.logger.Info("Incoming request: ",
			zap.String("ip address: ", r.RemoteAddr),
			zap.String("request endpoint: ", r.URL.Path),
			zap.String("request method: ", r.Method))

		req.ServeHTTP(rw, r)
	})
}
