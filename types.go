package fifteen

import (
	"github.com/yeoffrey/fifteen/places"
	"github.com/yeoffrey/fifteen/routing"
)

type LocationType string

const (
	Cafe   LocationType = "cafe"
	School LocationType = "school"
)

type FifteenRequest struct {
	OriginAddress string `json:"origin"`
}

type PlaceRoute struct {
	Place *places.Place  `json:"place"`
	Route *routing.Route `json:"route"`
}
