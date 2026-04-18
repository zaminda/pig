package tui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type updateAvailableMsg string

const latestReleaseURL = "https://api.github.com/repos/zaminda/pig/releases/latest"

func CheckLatestVersion(current string) tea.Cmd {
	return func() tea.Msg {
		latest, err := fetchLatestVersion()
		if err != nil || latest == "" {
			return nil
		}
		if normalise(latest) != normalise(current) && current != "dev" {
			return updateAvailableMsg(fmt.Sprintf("update available: %s", latest))
		}
		return nil
	}
}

func fetchLatestVersion() (string, error) {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(latestReleaseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}

func normalise(v string) string {
	return strings.TrimPrefix(v, "v")
}
