package main

import (
	"dot/client"
)

func main() {
	settings, err := client.GetSettings("settings.json")
	if err != nil {
		settings.WriteSettings()
	}

	locale, err := client.GetLocalization(settings.Language)
	if err != nil {
		panic(err)
	}

	InitializeWindow(settings.Width, settings.Height, settings.Fullscreen)

	rltextures, err := PreloadTextures()
	if err != nil {
		panic(err)
	}

	rlfonts, err := PreloadFonts()
	if err != nil {
		panic(err)
	}

	renderState := RenderState{
		settings:    settings,
		locale:      locale,
		rltextures:  rltextures,
		rlfonts:     rlfonts,
		uistate:     client.GetUIState(),
		buttonstate: GenButtonState(locale, rlfonts),
		camera:      InitCamera(),
	}

	RenderLoop(renderState)
}
