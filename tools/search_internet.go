package tools

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type tavilyRequest struct {
	ApiKey string `json:"api_key"`
	Query  string `json:"query"`
}

type tavilyResponse struct {
	Result []struct {
		Content string `json:"content"`
	} `json:"results"`
}

func SearchInternet(args map[string]string) string {
	query := args["query"]

	reqBody := tavilyRequest{
		ApiKey: os.Getenv("TAVILY_API_KEY"),
		Query:  query,
	}
	jsonData, _ := json.Marshal(reqBody)

	resp, err := http.Post("https://api.tavily.com/search", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var result tavilyResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if len(result.Result) == 0 {
		return "No result found"
	}

	if resp.StatusCode != http.StatusOK {
		return "Search failed: " + resp.Status
	}

	return result.Result[0].Content
}
