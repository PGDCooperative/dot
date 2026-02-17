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
		Width:      1280,
		Height:     720,
		Fullscreen: false,
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
