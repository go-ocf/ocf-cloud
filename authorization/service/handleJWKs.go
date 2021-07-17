package service

import (
	"github.com/valyala/fasthttp"
)

func (s *Service) HandleJWKs(ctx *fasthttp.RequestCtx) {
	if p, ok := s.deviceProvider.(interface {
		HandleJWKs(ctx *fasthttp.RequestCtx)
	}); ok {
		p.HandleJWKs(ctx)
		return
	}
	setErrorResponse(&ctx.Response, fasthttp.StatusNotFound, "not found")
}
