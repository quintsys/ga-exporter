package ga

import (
	"context"
	"log"

	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/googleapi"
)

// ViewData pulls data from a Google Analytics view
func ViewData(dimensions string) (*analytics.GaData, error) {
	ctx := context.Background()
	analyticsService, err := analytics.NewService(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: load view id and date paramaters from a config file or env
	dataGaGetCall := analyticsService.Data.Ga.
		Get("ga:115848806", "3daysAgo", "yesterday", "ga:sessions").
		Dimensions(dimensions)

	gaData, err := dataGaGetCall.Do()
	if err != nil {
		if gErr, ok := err.(*googleapi.Error); ok {
			log.Println("Error:", gErr.Code, "--", gErr.Message)
		}
		return nil, err
	}

	return gaData, nil
}
