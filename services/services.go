package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gabrielalmir/witch/models"
	"golang.design/x/clipboard"
)

func CopyToClipboard(s string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}

	clipboard.Write(clipboard.FmtText, []byte(s))
	return nil
}

func SearchRepos(query string) ([]models.Repo, error) {
	apiURL := fmt.Sprintf("https://api.github.com/search/repositories?q=%s+fork:false&sort=stars&order=desc&per_page=50", url.QueryEscape(query))

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Items []models.Repo `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Items, nil
}
