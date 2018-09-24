package Models

import (
	"github.com/bxcodec/faker"
)

type GameFilterOptionsModel struct {
	Id          string  `query:"id"`
	SteamAppId  string  `query:"appId"`
	Name        string  `query:"name"`
	Platforms   string  `query:"platforms"`
	Playability float32 `query:"playability"`
	SortBy      string  `query:"sort"`
	Page        int     `query:"page"`
	Limit       int     `query:"limit"`
}

type GameResultDataModel struct {
	Id            string `faker:unix_time`
	SteamAppId    string `faker:unix_time`
	Name          string `faker:name`
	ResultImage   string `faker:url`
	Platforms     []string
	CommentsCount int
	Playability   float32 `faker:`
	Description   string  `faker:sentence`
}

// Makes calls to postgres (or in this case our placeholder data) to get the results.
func FetchGamesByOptions(o GameFilterOptionsModel) (results []GameResultDataModel, err error) {
	// TODO: Add Postgres SQL statement.
	results, err = generatePlaceholderResults(o.Limit)
	return
}

// Creates placeholder data for our front end.
func generatePlaceholderResults(rows int) (results []GameResultDataModel, err error) {
	for i := 0; i < rows; i++ {
		// Generate mockup row.
		r := GameResultDataModel{}
		err = faker.FakeData(r)

		// If any errors we return.
		if err != nil {
			return
		}

		// If no errors we add the row to the results.
		results = append(results, r)
	}
	return
}
