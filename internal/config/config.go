package config

import (
	"os"
	"task-time-logger-go/internal/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	DBFilename   string
	JiraBaseURL  string
	JiraUsername string
	JiraAPIToken string
}

var AppConfig Config

func Load() {
	if err := godotenv.Load(); err != nil {
		logger.AppLogger.Fatalf("Couldn't Load configuration: %v", err)
	}

	AppConfig = Config{
		DBFilename:   "internal/storage/" + os.Getenv("DB_FILENAME"),
		JiraBaseURL:  os.Getenv("JIRA_BASE_URL"),
		JiraUsername: os.Getenv("JIRA_USERNAME"),
		JiraAPIToken: os.Getenv("JIRA_API_TOKEN"),
	}
}
