package ga

import (
	"context"
	"log"

	"github.com/quintsys/ga-exporter/env"
	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

// ViewData pulls data from a Google Analytics view
func ViewData(dimensions string) (*analytics.GaData, error) {
	ctx := context.Background()
	json := env.Lookup("GOOGLE_APPLICATION_CREDENTIALS", "{}")
	analyticsService, err := analytics.NewService(ctx, option.WithCredentialsJSON([]byte(json)))
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
