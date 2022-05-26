package middleware

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/shubhamjagdhane/simple-load-balancer/constant"
	"github.com/shubhamjagdhane/simple-load-balancer/errors"
	"github.com/shubhamjagdhane/simple-load-balancer/logger"
)

type ErrHandler func(http.ResponseWriter, *http.Request) error

func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			var stackSize int = 4 << 10 // 4 KB
			stack := make([]byte, stackSize)
			length := runtime.Stack(stack, true)
			logger.Log.Error(string(stack[:length]))
			xerr := errors.GetError("Server timeout", http.StatusRequestTimeout)
			w.WriteHeader(xerr.HTTPCode)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		w.Header().Set("Content-Type", constant.MIMEApplicationJson)
		xerr, ok := err.(*errors.CustomError)
		if !ok {
			xerr = errors.GetError("Server Unavailable", http.StatusServiceUnavailable)
		}

		w.WriteHeader(xerr.HTTPCode)
		_ = json.NewEncoder(w).Encode(xerr)
		return
	}
}
