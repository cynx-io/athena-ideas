package exampleservice

import (
	"context"
	"github.com/cynx-io/athena-ideas/internal/model/response"
	core "github.com/cynx-io/cynx-core/proto/gen"
)

func (s *Service) HealthCheck(ctx context.Context, req *core.GenericRequest, resp *core.GenericResponse) error {
	response.Success(resp)
	return nil
}
