package fifteen

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"
	"github.com/yeoffrey/fifteen/places"
	"github.com/yeoffrey/fifteen/routing"
	"googlemaps.github.io/maps"
)

var client *maps.Client

func loadSecret() string {
	if envApiKey := os.Getenv("GOOGLE_MAPS_API"); envApiKey != "" {
		return envApiKey
	}

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return os.Getenv("GOOGLE_MAPS_API")
}

func init() {
	functions.HTTP("Fifteen", Entrypoint)

	c, err := maps.NewClient(maps.WithAPIKey(loadSecret()))
	if err != nil {
		log.Fatalf(": %v\n", err)
	}
	client = c
}

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	var reqBody FifteenRequest

	// Parse request to a FifteenRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	originAddress := places.GetOrigin(client, reqBody.OriginAddress)

	// TODO: We can write a function to build this
	placesMap := map[LocationType]*places.Place{
		Cafe:   places.GetPlace(client, originAddress.ID, string(Cafe), PlaceTypeLookup[Cafe]),
		School: places.GetPlace(client, originAddress.ID, string(School), PlaceTypeLookup[School]),
	}

	placeRoutes := handleRouteMatrix(*originAddress, placesMap)

	// Send response
	jsonResponse, err := json.Marshal(placeRoutes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func handleRouteMatrix(origin places.Place, places map[LocationType]*places.Place) map[LocationType]*PlaceRoute {
	placeIDs := []string{}

	for _, place := range places {
		placeIDs = append(placeIDs, place.ID)
	}

	routeMatrix := routing.GetRouteMatrix(client, origin.ID, placeIDs)

	placeRoutes := map[LocationType]*PlaceRoute{}

	i := 0
	for key := range places {
		// TODO: What happens if there is 0 rows?
		// TODO: What happens if there if Status != OK
		routeMetadata := routeMatrix.Rows[0].Elements[i]

		placeRoutes[key] = &PlaceRoute{
			Place: places[key],
			Route: &routing.Route{
				Distance: routeMetadata.Distance.Meters,
				Duration: routeMetadata.Duration.Seconds(),
				Mode:     routing.Biking,
			},
		}

		i++
	}

	return placeRoutes
}
