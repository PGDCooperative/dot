package main

import (
	"dot/client"
)

func main() {
	settings, err := client.GetSettings("settings.json")
	if err != nil {
		panic(err)
	}
	InitializeWindow(settings.Width, settings.Height, settings.Fullscreen)
	rlassets, err := Preload("assets")
	if err != nil {
		panic(err)
	}
	err = RenderLoop(rlassets, settings)
	if err != nil {
		panic(err)
	}
}
