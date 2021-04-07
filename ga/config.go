package ga

import "github.com/quintsys/ga-exporter/env"

type config struct {
	credentials []byte
	viewID      string
	startDate   string
	endDate     string
	metric      string
}

func newConfig() *config {
	c := config{metric: "ga:sessions"}
	c.credentials = []byte(env.Lookup("GOOGLE_APPLICATION_CREDENTIALS", "{}"))
	c.viewID = env.Lookup("GA_VIEW_ID", "")
	c.startDate = env.Lookup("GA_START_DATE", "3daysAgo")
	c.endDate = env.Lookup("GA_END_DATE", "today")

	return &c
}
