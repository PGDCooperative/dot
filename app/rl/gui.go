package main

import (
	"dot/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func MainMenu(uistate *client.UIState, locale client.Locale, font *rl.Font) {
	rl.DrawTextEx(*font, "Divide. Operate. Thrive.",
		rl.NewVector2(20.0, 170.0), 100.0, 1.0, rl.White)
}

func SettingsMenu(uistate *client.UIState, settings *client.Settings, locale client.Locale, font *rl.Font) {

}

func PauseMenu(uistate *client.UIState, locale client.Locale, font *rl.Font) {

}
