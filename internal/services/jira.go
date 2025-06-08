package services

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"task-time-logger-go/internal/config"

	"github.com/tidwall/gjson"
)

func CallJiraAPI(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(config.AppConfig.JiraUsername + ":" + config.AppConfig.JiraAPIToken))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("JIRA API error: %s", string(body))
	}

	return body, nil
}

func GetJiraProjects() ([]string, error) {
	url := config.AppConfig.JiraBaseURL + "/rest/api/3/project/search"
	body, err := CallJiraAPI(url)
	if err != nil {
		return nil, err
	}

	gjsonValue := gjson.Get(string(body), "values.#.key")
	var keys []string
	for _, item := range gjsonValue.Array() {
		keys = append(keys, item.String())
	}

	if len(keys) == 0 {
		return nil, fmt.Errorf("no projects found")
	}

	return keys, nil
}
