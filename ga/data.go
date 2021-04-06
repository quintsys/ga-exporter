package ga

// Data holds the information to be pulled from GA
type Data struct {
	ClientID       string `json:"client_id,omitempty"`
	Campaign       string `json:"campaign,omitempty"`
	Source         string `json:"source,omitempty"`
	Medium         string `json:"medium,omitempty"`
	Keyword        string `json:"keyword,omitempty"`
	AdGroup        string `json:"ad_group,omitempty"`
	AdContent      string `json:"ad_content,omitempty"`
	AdMatchedQuery string `json:"ad_matched_query,omitempty"`
}

// AdData pulls ad related data from GA
func AdData() ([]Data, error) {
	dimensions := "ga:dimension5,ga:adGroup,ga:adContent,ga:adMatchedQuery"
	gaData, err := ViewData(dimensions)
	if err != nil {
		return nil, err
	}

	data := []Data{}
	for _, row := range gaData.Rows {
		data = append(data, Data{
			ClientID:       row[0],
			AdGroup:        row[1],
			AdContent:      row[2],
			AdMatchedQuery: row[3],
		})
	}

	return data, nil
}

// OriginData pulls origin related data from GA
func OriginData() ([]Data, error) {
	dimensions := "ga:dimension5,ga:campaign,ga:source,ga:medium,ga:keyword"
	gaData, err := ViewData(dimensions)
	if err != nil {
		return nil, err
	}

	data := []Data{}
	for _, row := range gaData.Rows {
		data = append(data, Data{
			ClientID: row[0],
			Campaign: row[1],
			Source:   row[2],
			Medium:   row[3],
			Keyword:  row[4],
		})
	}

	return data, nil
}
