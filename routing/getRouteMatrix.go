package routing

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

func GetRouteMatrix(client *maps.Client, origin string, destination []string) *maps.DistanceMatrixResponse {
	distance, err := client.DistanceMatrix(context.Background(), &maps.DistanceMatrixRequest{
		Origins:      transform([]string{origin}),
		Destinations: transform(destination),
		Mode:         maps.TravelModeBicycling,
	})
	if err != nil {
		log.Print(err)
		return nil
	}

	return distance
}

func transform(arr []string) []string {
	for i := range arr {
		arr[i] = "place_id:" + arr[i]
	}

	return arr
}
