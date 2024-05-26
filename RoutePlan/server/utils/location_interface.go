package utils

import (
	"strconv"
	"strings"
)

type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
	Distance  float64
	Cost      float64
	Rating    float64
}

type Route struct {
	Locations []*Location
	Feature   RouteFeature
}

type RouteFeature struct {
	TotalCost     float64
	TotalRating   float64
	AverageCost   float64
	AverageRating float64
}

func FromPoi(poi Poi) (Location, error) {
	var (
		location Location
		err      error
	)
	location.Name = poi.Name
	location.Distance, err = strconv.ParseFloat(poi.Distance, 64)
	if err != nil {
		return location, err
	}

	gpsInfo := strings.Split(poi.Location, ",")
	location.Longitude, err = strconv.ParseFloat(gpsInfo[0], 64)
	if err != nil {
		return location, err
	}
	location.Latitude, err = strconv.ParseFloat(gpsInfo[1], 64)
	if err != nil {
		return location, err
	}

	if costStr, ok := poi.BusinessExt.Cost.(string); ok {
		if cost, err := strconv.ParseFloat(costStr, 64); err == nil {
			location.Cost = cost
		} else {
			return location, err
		}
	}

	if ratingStr, ok := poi.BusinessExt.Rating.(string); ok {
		if rating, err := strconv.ParseFloat(ratingStr, 64); err == nil {
			location.Rating = rating
		} else {
			return location, err
		}
	}

	return location, nil
}
