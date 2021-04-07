package ga

import (
	"context"
	"log"

	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

// ViewData pulls data from a Google Analytics view
func ViewData(dimensions string) (*analytics.GaData, error) {
	ctx := context.Background()
	cfg := newConfig()
	analyticsService, err := analytics.NewService(ctx, option.WithCredentialsJSON(cfg.credentials))
	if err != nil {
		return nil, err
	}

	dataGaGetCall := analyticsService.Data.Ga.
		Get(cfg.viewID, cfg.startDate, cfg.endDate, cfg.metric).
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
