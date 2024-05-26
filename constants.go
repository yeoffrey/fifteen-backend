package fifteen

import (
	"googlemaps.github.io/maps"
)

var PlaceTypeLookup = map[LocationType]maps.PlaceType{
	Cafe:   maps.PlaceTypeCafe,
	School: maps.PlaceTypeSchool,
}
