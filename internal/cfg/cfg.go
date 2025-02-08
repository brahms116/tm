package cfg

import (
	"os"
)

type Cfg struct {
	DbUrl  string
	ApiKey string
}

func Must() Cfg {
	dbUrl := os.Getenv("TM_DB_URL")
	if dbUrl == "" {
		panic("TM_DB_URL must be set")
	}

	apiKey := os.Getenv("TM_API_KEY")
	if apiKey == "" {
		panic("TM_API_KEY must be set")
	}

	return Cfg{dbUrl, apiKey}
}
