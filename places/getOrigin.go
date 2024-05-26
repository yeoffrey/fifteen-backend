package places

import (
	"context"

	"googlemaps.github.io/maps"
)

func GetOrigin(client *maps.Client, origin string) *Place {
	ctx := context.WithValue(context.Background(), "X-Goog-FieldMask", FIELD_MASK)

	res, err := client.TextSearch(ctx, &maps.TextSearchRequest{
		Query: origin,
	})
	if err != nil {
		return nil
	}

	// TODO: What happens if there are no results from Google?

	result := res.Results[0]

	return &Place{
		ID:      result.PlaceID,
		Name:    result.Name,
		Address: result.FormattedAddress,
	}
}
