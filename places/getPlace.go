package places

import (
	"context"

	"googlemaps.github.io/maps"
)

func GetPlace(client *maps.Client, origin string, locationType string, placeType maps.PlaceType) *Place {
	ctx := context.WithValue(context.Background(), "X-Goog-FieldMask", FIELD_MASK)

	res, err := client.TextSearch(ctx, &maps.TextSearchRequest{
		Query: locationType + " near " + origin,
		Type:  placeType,
	})
	if err != nil {
		return nil
	}

	// TODO: What happens if there are no results from Google?

	result := res.Results[0]

	info := &Place{
		ID:      result.PlaceID,
		Name:    result.Name,
		Address: result.FormattedAddress,
	}

	return info
}
