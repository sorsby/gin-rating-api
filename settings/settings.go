package settings

// APISettings stores all the environent variables for the API.
type APISettings struct {
	GinRatingTableName string `env:"GIN_RATING_TABLE_NAME"`
	CognitoRegion      string `env:"COGNITO_REGION"`
	CognitoUserPool    string `env:"COGNITO_USER_POOL"`
}
