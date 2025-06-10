package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBFilename   string
	JiraBaseURL  string
	JiraUsername string
	JiraAPIToken string
}

var AppConfig Config

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	AppConfig = Config{
		DBFilename:   "internal/storage/" + os.Getenv("DB_FILENAME"),
		JiraBaseURL:  os.Getenv("JIRA_BASE_URL"),
		JiraUsername: os.Getenv("JIRA_USERNAME"),
		JiraAPIToken: os.Getenv("JIRA_API_TOKEN"),
	}

	return nil
}
