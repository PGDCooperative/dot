package main

import (
	"dot/client"
)

func main() {
	settings, err := client.GetSettings("settings.json")
	if err != nil {
		panic(err)
	}
	locale, err := client.GetLocalization(settings.Language)
	if err != nil {
		panic(err)
	}
	InitializeWindow(settings.Width, settings.Height, settings.Fullscreen)
	rlassets, err := Preload()
	if err != nil {
		panic(err)
	}
	err = RenderLoop(rlassets, settings, locale)
	if err != nil {
		panic(err)
	}
}
