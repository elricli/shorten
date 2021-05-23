package recovery

import (
	"net/http"

	"github.com/drrrMikado/shorten/pkg/middleware"
	"go.uber.org/zap"
)

// Option is recovery option.
type Option func(*option)

type option struct {
	logger *zap.SugaredLogger
}

// WithLogger with recovery logger.
func WithLogger(logger *zap.SugaredLogger) Option {
	return func(o *option) {
		o.logger = logger.Named("middleware.recovery")
	}
}

func Recovery(opts ...Option) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		var opt option
		for _, o := range opts {
			o(&opt)
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					opt.logger.Error("catch a error：", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
			return
		})
	}
}
