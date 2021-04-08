package ga

import (
	"context"
	"errors"
	"log"

	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

// ViewData pulls data from a Google Analytics view
// Do executes the "analytics.data.ga.get" call.
// Exactly one of *GaData or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *GaData.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func ViewData(dimensions string) (*analytics.GaData, error) {
	cfg := newConfig()
	if len(cfg.credentials) == 0 {
		return nil, errors.New("GOOGLE_APPLICATION_CREDENTIALS not set")
	}

	credentialsOption := option.WithCredentialsJSON(cfg.credentials)

	svc, err := analytics.NewService(context.Background(), credentialsOption)
	if err != nil {
		return nil, err
	}

	gaData, err := svc.Data.Ga.
		Get(cfg.viewID, cfg.startDate, cfg.endDate, cfg.metric).
		Dimensions(dimensions).
		Do()
	if err != nil {
		logGoogleError(err)
		return nil, err
	}

	return gaData, nil
}

func logGoogleError(err error) {
	if gErr, ok := err.(*googleapi.Error); ok {
		log.Println("Error:", gErr.Code, "--", gErr.Message)
		log.Println("Details:", gErr.Details)
	}
}
