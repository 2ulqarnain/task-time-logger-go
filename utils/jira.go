package utils

import (
	"io"
	"net/http"
	"task-time-logger-go/utils/vars"
)

func CallJiraApi(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(vars.JIRA_USERNAME, vars.JIRA_API_TOKEN)

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
