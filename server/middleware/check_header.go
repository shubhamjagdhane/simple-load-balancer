package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shubhamjagdhane/simple-load-balancer/constant"
	"github.com/shubhamjagdhane/simple-load-balancer/errors"
	"github.com/shubhamjagdhane/simple-load-balancer/tracer"
)

func ValidateHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx, span := tracer.StartSpan(req.Context(), "ValidateHeaderMiddleware")
		defer span.End()

		w.Header().Set(constant.HeaderContentType, constant.MIMEApplicationJson)

		if err := validateHeaders(ctx, req); err != nil {
			span.RecordError(err)

			w.WriteHeader(err.HTTPCode)
			_ = json.NewEncoder(w).Encode(err)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func validateHeaders(ctx context.Context, req *http.Request) *errors.CustomError {
	if req == nil {
		return errors.GetError(errors.BadRequest, http.StatusBadRequest)
	}

	if !isEqual(req, constant.HeaderContentType, constant.MIMEApplicationJson) {
		return errors.GetError(errors.BadContentType, http.StatusBadRequest)
	}
	if !isEqual(req, constant.HeaderAccept, constant.MIMEApplicationJson) {
		return errors.GetError(errors.BadHeader, http.StatusBadRequest)
	}
	return nil
}

func isEqual(req *http.Request, check, value string) bool {
	if req == nil {
		return false
	}
	return req.Header.Get(check) == value
}
