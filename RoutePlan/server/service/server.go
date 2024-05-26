package service

import (
	servicepb "route_plan/server/proto/service"
	"route_plan/server/utils"
)

type Server struct {
	servicepb.UnimplementedRoutePlanServer

	GaodeClient utils.GaodeHttpClientInterface
}
