package utils

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CallJiraApi(url string) ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, Error:%s", err.Error())
	}

	username := os.Getenv("JIRA_USERNAME")
	password := os.Getenv("JIRA_API_TOKEN")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
