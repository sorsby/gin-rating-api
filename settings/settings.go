package settings

// APISettings stores all the environent variables for the API.
type APISettings struct {
	GinRatingTableName string `env:"GIN_RATING_TABLE_NAME"`
}
