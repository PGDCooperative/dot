package client

import (
	"encoding/json"
	"os"
)

type Settings struct {
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Fullscreen bool   `json:"fullscreen"`
	Language   string `json:"language"`
}

func GetSettings(path string) (Settings, error) {
	// default value
	settings := Settings{
		Width:      1920,
		Height:     1080,
		Fullscreen: true,
		Language:   "english",
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return settings, err
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return settings, err
	}
	return settings, nil
}

func (s Settings) WriteSettings() error {
	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("settings.json", data, 0770)
	if err != nil {
		return err
	}
	return nil
}
