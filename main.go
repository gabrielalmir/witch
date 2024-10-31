package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gabrielalmir/witch/models"
	"github.com/gabrielalmir/witch/services"
)

type model struct {
	repos        []models.Repo
	query        string
	status       string
	cursor       int
	page         int
	itemsPerPage int
	selectedRepo models.Repo
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.query == "" {
				link := fmt.Sprintf("github.com/%s/%s", m.selectedRepo.Owner.Login, m.selectedRepo.Name)
				err := services.CopyToClipboard(link)
				if err != nil {
					m.status = "Error copying to clipboard: " + err.Error()
				} else {
					m.status = "GitHub link copied to clipboard!"
				}
			} else {
				repos, err := services.SearchRepos(m.query)
				if err != nil {
					m.status = "Error fetching repositories: " + err.Error()
				} else {
					m.repos = repos
					m.status = fmt.Sprintf("Results for: %s", m.query)
					m.query = ""
					m.cursor = 0
					m.page = 0
				}
			}
		case "ctrl+c", "q", "esc":
			fmt.Println("Happy Halloween! 🎃")
			return m, tea.Quit
		case "backspace":
			if len(m.query) > 0 {
				m.query = m.query[:len(m.query)-1]
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else if m.page > 0 {
				m.page--
				m.cursor = m.itemsPerPage - 1
			}
		case "down":
			if m.cursor < m.itemsPerPage-1 && m.cursor < len(m.getCurrentPageItems())-1 {
				m.cursor++
			} else if m.page < (len(m.repos)-1)/m.itemsPerPage {
				// Go to the next page
				m.page++
				m.cursor = 0
			}
		case "left":
			if m.page > 0 {
				m.page--
			}
		case "right":
			if m.page < (len(m.repos)-1)/m.itemsPerPage {
				m.page++
			}
		default:
			m.query += msg.String()
		}
	}

	if m.getCurrentPageItems() != nil && len(m.getCurrentPageItems()) > 0 {
		m.selectedRepo = m.getCurrentPageItems()[m.cursor]
	}

	return m, nil
}

func (m model) View() string {
	s := "Welcome to Witch! 🧙‍♀️\n\n"
	s += "Find 'witch' repositories are trending on GitHub!\n"
	s += "Enter a keyword and press Enter:\n\n"
	s += fmt.Sprintf("Current search: %s\n\n", m.query)
	s += m.status + "\n\n"

	pageItems := m.getCurrentPageItems()
	for i, repo := range pageItems {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s/%s (★ %d)\n  %s\n\n", cursor, repo.Owner.Login, repo.Name, repo.Stars, repo.Description)
	}

	s += fmt.Sprintf("Page %d/%d\n", m.page+1, (len(m.repos)-1)/m.itemsPerPage+1)

	return s
}

func (m model) getCurrentPageItems() []models.Repo {
	start := m.page * m.itemsPerPage
	end := start + m.itemsPerPage
	if end > len(m.repos) {
		end = len(m.repos)
	}
	return m.repos[start:end]
}

func main() {
	m := model{
		status:       "Type a keyword and press Enter to search.",
		itemsPerPage: 5,
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
