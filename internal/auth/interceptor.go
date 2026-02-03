package auth

import (
	"context"
	"fmt"
	"strings"

	"connectrpc.com/connect"
)

type Interceptor struct {
	validator *Auth0Validator
	skip      map[string]struct{}
}

func NewInterceptor(validator *Auth0Validator, skipProcedures []string) *Interceptor {
	skip := map[string]struct{}{}
	for _, p := range skipProcedures {
		skip[p] = struct{}{}
	}
	return &Interceptor{
		validator: validator,
		skip:      skip,
	}
}

func (i *Interceptor) Unary() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if _, ok := i.skip[req.Spec().Procedure]; ok {
				return next(ctx, req)
			}

			authHeader := req.Header().Get("Authorization")
			if authHeader == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("missing authorization header"))
			}
			token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid authorization header"))
			}

			if _, err := i.validator.Validate(ctx, token); err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			return next(ctx, req)
		}
	}
}
