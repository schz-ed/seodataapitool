package dataforseo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type KeywordMetrics struct {
	Keyword    string  `json:"keyword"`
	Volume     int     `json:"volume"`
	Difficulty float64 `json:"difficulty"`
}

// FetchKeywordData simulates a call to the DataForSEO API.
// In a real implementation, build the request according to API documentation.
func FetchKeywordData(keywords []string, apiKey string) ([]KeywordMetrics, error) {
	reqBody, err := json.Marshal(map[string]interface{}{
		"keywords": keywords,
	})
	if err != nil {
		return nil, err
	}

	// Replace with the actual DataForSEO API endpoint.
	url := "https://api.dataforseo.com/v3/keywords_data/keyword_overview/live"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(bodyBytes))
	}

	var metrics []KeywordMetrics
	err = json.NewDecoder(resp.Body).Decode(&metrics)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}
