package models

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Owner       struct {
		Login string `json:"login"`
	} `json:"owner"`
}
