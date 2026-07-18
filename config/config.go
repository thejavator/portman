package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// AppConfig represents the application configuration
type AppConfig struct {
	Favorites []string `json:"favorites"`
	Theme     string   `json:"theme"`
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(home, ".config", "portman")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// LoadConfig loads the configuration from ~/.config/portman/config.json
func LoadConfig() AppConfig {
	cfg := AppConfig{
		Favorites: []string{"docker", "postgres", "redis", "mysql", "launchd"}, // Default
		Theme:     "default",
	}
	path, err := getConfigPath()
	if err != nil {
		return cfg
	}

	data, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(data, &cfg)
	}
	return cfg
}

// SaveConfig saves the configuration to the JSON file
func SaveConfig(cfg AppConfig) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// IsFavorite checks if a process is in the favorites list
func (c *AppConfig) IsFavorite(procName string) bool {
	for _, f := range c.Favorites {
		if f == procName {
			return true
		}
	}
	return false
}

// ToggleFavorite adds or removes a process from the favorites list
func (c *AppConfig) ToggleFavorite(procName string) {
	for i, f := range c.Favorites {
		if f == procName {
			// Remove it
			c.Favorites = append(c.Favorites[:i], c.Favorites[i+1:]...)
			SaveConfig(*c)
			return
		}
	}
	// Add it
	c.Favorites = append(c.Favorites, procName)
	SaveConfig(*c)
}
