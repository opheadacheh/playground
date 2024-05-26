package service

import (
	"context"
	"fmt"
	"log"
	"route_plan/server/utils"

	commonpb "route_plan/server/proto/common"
	routepb "route_plan/server/proto/route"
	servicepb "route_plan/server/proto/service"
)

func generateRoutes(locationList [][]utils.Location, routes *[]utils.Route, currentRoute utils.Route) {
	if len(locationList) == 0 {
		newStopPoints := make([]*utils.Location, len(currentRoute.Locations))
		_ = copy(newStopPoints, currentRoute.Locations)
		currentRoute.Locations = newStopPoints
		*routes = append(*routes, currentRoute)
		return
	}

	for i := range locationList[0] {
		currentRoute.Locations = append(currentRoute.Locations, &locationList[0][i])
		generateRoutes(locationList[1:], routes, currentRoute)
		currentRoute.Locations = currentRoute.Locations[0 : len(currentRoute.Locations)-1]
	}
}

func convertRouteToRouteInfo(route *utils.Route, tag routepb.RouteTag) (*routepb.RouteInfo, error) {
	routeInfo := &routepb.RouteInfo{
		RouteTag: tag,
	}
	for _, location := range route.Locations {
		routeInfo.StopPoints = append(routeInfo.StopPoints, &routepb.StopPoint{
			Name:   location.Name,
			Rating: fmt.Sprintf("%.2f", location.Rating),
			Cost:   fmt.Sprintf("%.2f", location.Cost),
			GpsInfo: &commonpb.GpsInfo{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			},
		})
	}

	return routeInfo, nil
}

func getLeaseCostRoute(routes []utils.Route) (*utils.Route, error) {
	var (
		minCost        = float64(10000000)
		leastCostRoute *utils.Route
	)

	for i := range routes {
		if routes[i].Feature.AverageCost < minCost {
			minCost = routes[i].Feature.AverageCost
			leastCostRoute = &routes[i]
		}
	}

	return leastCostRoute, nil
}

func getHighestRatingRoute(routes []utils.Route) (*utils.Route, error) {
	var (
		maxRating          = float64(0)
		highestRatingRoute *utils.Route
	)

	for i := range routes {
		if routes[i].Feature.AverageRating > maxRating {
			maxRating = routes[i].Feature.AverageRating
			highestRatingRoute = &routes[i]
		}
	}

	return highestRatingRoute, nil
}

func populateRouteFeature(routes []utils.Route) error {
	for i := range routes {
		var costNum, ratingNum int
		for _, location := range routes[i].Locations {
			if location.Cost != 0 {
				routes[i].Feature.TotalCost += location.Cost
				costNum += 1
			}

			if location.Rating != 0 {
				routes[i].Feature.TotalRating += location.Rating
				ratingNum += 1
			}
		}

		if costNum > 0 {
			routes[i].Feature.AverageCost = routes[i].Feature.TotalCost / float64(costNum)
		} else {
			routes[i].Feature.AverageCost = 999999
		}

		if ratingNum > 0 {
			routes[i].Feature.AverageRating = routes[i].Feature.TotalRating / float64(ratingNum)
		} else {
			routes[i].Feature.AverageRating = 0.0001
		}
	}

	return nil
}

func (s *Server) getLocationList(keywords []string, gpsInfo *commonpb.GpsInfo) ([][]utils.Location, error) {
	var locationsList [][]utils.Location
	for _, keyword := range keywords {
		if locations, err := s.GaodeClient.Search(gpsInfo, keyword, 10); err != nil {
			log.Printf("Search got error: %v", err)
			return nil, err
		} else {
			locationsList = append(locationsList, locations)
		}
	}

	return locationsList, nil
}

func (s *Server) PlanRoute(ctx context.Context, in *servicepb.PlanRouteRequest) (*servicepb.PlanRouteResponse, error) {
	log.Printf("got request: %v\n", in)

	locationList, err := s.getLocationList(in.GetKeywords(), in.GetGpsInfo())
	if err != nil {
		return nil, err
	}

	if len(locationList) == 0 {
		return &servicepb.PlanRouteResponse{}, nil
	}

	var (
		routes []utils.Route
		route  utils.Route
	)

	route.Locations = make([]*utils.Location, 0, len(locationList))
	generateRoutes(locationList, &routes, route)

	populateRouteFeature(routes)

	for _, route := range routes {
		log.Println(route)
	}

	var response servicepb.PlanRouteResponse

	leastCostRoute, err := getLeaseCostRoute(routes)
	if err != nil {
		return nil, err
	}

	fmt.Println(leastCostRoute)

	leastCostRouteInfo, err := convertRouteToRouteInfo(leastCostRoute, routepb.RouteTag_LEAST_EXPENSIVE)
	if err != nil {
		return nil, err
	}
	response.RouteInfos = append(response.RouteInfos, leastCostRouteInfo)

	highestRatingRoute, err := getHighestRatingRoute(routes)
	if err != nil {
		return nil, err
	}

	fmt.Println(highestRatingRoute)

	highestRatingRouteInfo, err := convertRouteToRouteInfo(highestRatingRoute, routepb.RouteTag_HIGHEST_RATING)
	if err != nil {
		return nil, err
	}
	response.RouteInfos = append(response.RouteInfos, highestRatingRouteInfo)

	return &response, nil
}
