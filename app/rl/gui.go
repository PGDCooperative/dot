package main

import (
	"dot/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func MainMenu(uistate *client.UIState, buttonstate ButtonState, font rl.Font) {
	rl.DrawTextEx(font, "Divide. Operate. Thrive.",
		rl.NewVector2(20.0, 150.0), 100.0, 1.0, rl.White)

	mousePos := rl.GetMousePosition()
	buttonstate["NewGame"].Draw(mousePos, func() {

	})
	buttonstate["LoadGame"].Draw(mousePos, func() {

	})
	buttonstate["Settings"].Draw(mousePos, func() {

	})
	buttonstate["QuitGame"].Draw(mousePos, func() {

	})
}

func SettingsMenu(uistate *client.UIState, buttonstate ButtonState, settings *client.Settings) {

}

func PauseMenu(uistate *client.UIState, buttonstate ButtonState) {

}
